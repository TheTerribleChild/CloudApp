package cacheutil

import (
	"time"
)

type CacheClient interface {
	Set(key string, value interface{}, ttl int) error
	SetJsonCompress(key string, value interface{}, ttl int) error 
	Get(key string) (interface{}, error) 
	GetString(key string) (string, error)
	GetJsonDecompress(key string, item interface{}) error 
	Delete(keys ...interface{}) (int, error) 
	Expire(key string, ttl int) error
	GetCurrentTime() (time.Time, error) 
}