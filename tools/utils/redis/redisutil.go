package redisutil

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"time"
	"github.com/gomodule/redigo/redis"
)

const (
	PING = "PING"
	SET  = "SET"
	GET  = "GET"
	DEL  = "DEL"
	TIME = "TIME"
)

type RedisClient struct {
	pool *redis.Pool
}

type RedisClientBuilder struct {
	url                 string
	password            string
	maxActiveConnection int
	maxIdleConnection   int
}

func GetRedisClientBuilder(url string) *RedisClientBuilder {
	return &RedisClientBuilder{url: url}
}

func (instance *RedisClientBuilder) SetPassword(password string) *RedisClientBuilder {
	instance.password = password
	return instance
}

func (instance *RedisClientBuilder) SetMaxActiveConnection(maxActiveConnection int) *RedisClientBuilder {
	instance.maxActiveConnection = maxActiveConnection
	return instance
}

func (instance *RedisClientBuilder) SetMaxIdleConnection(maxIdleConnection int) *RedisClientBuilder {
	instance.maxIdleConnection = maxIdleConnection
	return instance
}

func (instance *RedisClientBuilder) Build() (*RedisClient, error) {
	options := make([]redis.DialOption, 0)
	if len(instance.url) == 0 {
		instance.url = ":6379"
	}
	if len(instance.password) > 0 {
		options = append(options, redis.DialPassword(instance.password))
	}
	if instance.maxActiveConnection == 0 {
		instance.maxActiveConnection = 1000
	}
	if instance.maxIdleConnection == 0 {
		instance.maxIdleConnection = 10
	}

	pool := &redis.Pool{
		MaxActive: instance.maxActiveConnection,
		MaxIdle:   instance.maxIdleConnection,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", instance.url, options...)
			if err != nil {
				return c, err
			}
			return c, err
		},
	}
	conn := pool.Get()
	defer conn.Close()
	return &RedisClient{pool: pool}, Ping(conn)
}

func (instance *RedisClient) GetConnection() redis.Conn {
	return instance.pool.Get()
}

func (instance *RedisClient) Ping() error {
	conn := instance.pool.Get()
	defer conn.Close()
	return Ping(conn)
}

func (instance *RedisClient) Set(key string, value interface{}) error {
	conn := instance.pool.Get()
	defer conn.Close()
	return Set(conn, key, value)
}

func (instance *RedisClient) SetJsonCompress(key string, value interface{}) error {
	conn := instance.pool.Get()
	defer conn.Close()
	return SetJsonCompress(conn, key, value)
}

func (instance *RedisClient) Get(key string) (interface{}, error) {
	conn := instance.pool.Get()
	defer conn.Close()
	return Get(conn, key)
}

func (instance *RedisClient) GetString(key string) (string, error) {
	conn := instance.pool.Get()
	defer conn.Close()
	return redis.String(Get(conn, key))
}

func (instance *RedisClient) GetJsonDecompress(key string, item interface{}) error {
	conn := instance.pool.Get()
	defer conn.Close()
	return GetJsonDecompress(conn, key, item)
}

func (instance *RedisClient) Delete(keys ...interface{}) (int, error) {
	conn := instance.pool.Get()
	defer conn.Close()
	return Delete(conn, keys...)
}

func (instance *RedisClient) GetCurrentTime() (time.Time, error) {
	conn := instance.pool.Get()
	defer conn.Close()
	return GetCurrentTime(conn)
}

//Generic Redis function

func Ping(conn redis.Conn) error {
	pong, err := conn.Do(PING)
	if err != nil {
		return err
	}

	_, err = redis.String(pong, err)
	if err != nil {
		return err
	}

	return nil
}

func Set(conn redis.Conn, key string, value interface{}) error {
	_, err := conn.Do(SET, key, value)
	return err
}

func Get(conn redis.Conn, key string) (interface{}, error) {
	return conn.Do(GET, key)
}

func Delete(conn redis.Conn, keys ...interface{}) (int, error) {
	return redis.Int(conn.Do(DEL, keys...))
}

func SetJsonCompress(conn redis.Conn, key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	writer.Write(b)
	writer.Flush()
	writer.Close()
	return Set(conn, key, buffer.Bytes())
}

func GetJsonDecompress(conn redis.Conn, key string, item interface{}) error {
	value, err := Get(conn, key)
	if err != nil {
		return err
	}
	if value == nil {
		return redis.ErrNil
	}
	reader, err := gzip.NewReader(bytes.NewReader(value.([]byte)))
	defer reader.Close()
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	json.Unmarshal(data, item)
	return nil
}

func GetCurrentTime(conn redis.Conn) (time.Time, error){
	times, err := redis.Int64s(conn.Do(TIME))
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(times[0], times[1]), nil
}

func IsEmptyError(err error) bool {
	if err == redis.ErrNil {
		return true
	}
	return false
}