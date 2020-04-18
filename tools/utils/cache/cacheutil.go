package cacheutil

import (
	"fmt"
)

type ICacheClient interface {
	Set(key string, value interface{}, ttl int) error
	Get(key string) (interface{}, error)
	GetString(key string) (string, error)
	StoreObject(key string, value interface{}, ttl int) error
	ScanObject(key string, item interface{}) error
	Delete(keys ...interface{}) (int, error) 
	Expire(key string, ttl int) error
}

type ICacheKey interface {
	GetKeyString() string
}

type CacheKey struct {
	Key string
}

func (instance CacheKey) GetKeyString() string {
	return instance.Key
}

type LoadingCache struct {
	CacheClient     ICacheClient
	LoadingFunction func(ICacheKey) (interface{}, error)

	TTL int

	Service string
	Collection string
}

func (instance *LoadingCache) Get(key ICacheKey) (result interface{}, err error) {

	cacheKey := instance.getKeyString(key)
	if result, err = instance.CacheClient.Get(cacheKey); err != nil || result != nil {
		return
	}
	if result, err = instance.LoadingFunction(key); err != nil || result == nil {
		return
	}
	instance.CacheClient.Set(cacheKey, result, instance.TTL)
	return
}

func (instance *LoadingCache) getKeyString(key ICacheKey) string {
	return fmt.Sprintf("Service:%s,Collection:%s,ID:%s", instance.Service, instance.Collection, key.GetKeyString())
}