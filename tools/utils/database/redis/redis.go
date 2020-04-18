package redisutil

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gomodule/redigo/redis"
	"theterriblechild/CloudApp/tools/utils/database/databaseconfig"
)

const (
	PING    = "PING"
	SET     = "SET"
	GET     = "GET"
	MGET    = "MGET"
	DEL     = "DEL"
	TIME    = "TIME"
	EXPIRE  = "EXPIRE"
	MULTI   = "MULTI"
	EXEC    = "EXEC"
	DISCARD = "DISCARD"
)

type RedisClient struct {
	pool *redis.Pool
}

type RedisClientBuilder struct {
	Host                string
	Password            string
	Port                int
	MaxActiveConnection int
	MaxIdleConnection   int
}

func GetRedisClient(config databaseconfig.DatabaseConfig) (*RedisClient, error) {
	options := make([]redis.DialOption, 0)
	if config.Port == 0 {
		config.Port = 6379
	}
	connectionString := fmt.Sprintf("%s:%d", config.Host, config.Port)
	if len(config.Password) > 0 {
		options = append(options, redis.DialPassword(config.Password))
	}
	if config.MaxConns == 0 {
		config.MaxConns = 1000
	}
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 10
	}

	pool := &redis.Pool{
		MaxActive: config.MaxConns,
		MaxIdle:   config.MaxIdleConns,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", connectionString, options...)
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

func (instance *RedisClient) StoreObject(key string, value interface{}, ttl int) error {
	conn := instance.pool.Get()
	defer conn.Close()
	if ttl > 0 {
		BeginTxn(conn)
		StoreObject(conn, key, value)
		Expire(conn, key, ttl)
		CommitTxn(conn)
		if err := conn.Flush(); err != nil {
			DiscardTxn(conn)
			return err
		}
		return nil
	} else {
		Set(conn, key, value)
		if err := conn.Flush(); err != nil {
			DiscardTxn(conn)
			return err
		}
		return nil
	}
}

func (instance *RedisClient) ScanObject(key string, item interface{}) error {
	conn := instance.pool.Get()
	defer conn.Close()
	return ScanObject(conn, key, item)
}

func (instance *RedisClient) Set(key string, value interface{}, ttl int) error {
	conn := instance.pool.Get()
	defer conn.Close()
	if ttl > 0 {
		BeginTxn(conn)
		Set(conn, key, value)
		Expire(conn, key, ttl)
		CommitTxn(conn)
		if err := conn.Flush(); err != nil {
			DiscardTxn(conn)
			return err
		}
		return nil
	} else {
		Set(conn, key, value)
		if err := conn.Flush(); err != nil {
			DiscardTxn(conn)
			return err
		}
		return nil
	}
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

func (instance *RedisClient) Delete(keys ...interface{}) (int, error) {
	conn := instance.pool.Get()
	defer conn.Close()
	return Delete(conn, keys...)
}

func (instance *RedisClient) Expire(key string, ttl int) error {
	conn := instance.pool.Get()
	defer conn.Close()
	return Expire(conn, key, ttl)
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
	return conn.Send(SET, key, value)
}

func Get(conn redis.Conn, key string) (interface{}, error) {
	return conn.Do(GET, key)
}

//func Scan(conn redis.Conn, key string, response interface{}) error {
//	reply, err := redis.Values(conn.Do(MGET, key))
//	if err != nil {
//		return err
//	}
//	if _, err := redis.Scan(reply, response); err != nil {
//		return err
//	}
//	return nil
//}

func Delete(conn redis.Conn, keys ...interface{}) (int, error) {
	return redis.Int(conn.Do(DEL, keys...))
}

func Expire(conn redis.Conn, key string, ttl int) error {
	return conn.Send(EXPIRE, key, ttl)
}

func GetCurrentTime(conn redis.Conn) (time.Time, error) {
	times, err := redis.Int64s(conn.Do(TIME))
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(times[0], times[1]), nil
}

func BeginTxn(conn redis.Conn) error {
	return conn.Send(MULTI)
}

func CommitTxn(conn redis.Conn) error {
	return conn.Send(EXEC)
}

func DiscardTxn(conn redis.Conn) error {
	return conn.Send(DISCARD)
}

func IsEmptyError(err error) bool {
	if err == redis.ErrNil {
		return true
	}
	return false
}

func StoreObject(conn redis.Conn, key string, value interface{}) error {
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

func ScanObject(conn redis.Conn, key string, item interface{}) error {
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
