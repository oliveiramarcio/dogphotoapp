package main

import (
	"sync"
	"time"
)

type CacheRepository interface {
	Get(breed string) (string, bool)
	Set(breed, url string)
}

type CachedValue struct {
	value      string
	expiration time.Time
}

type InMemoryCache struct {
	cache    sync.Map
	duration time.Duration
}

func (c *InMemoryCache) Get(breed string) (string, bool) {
	value, ok := c.cache.Load(breed)
	if !ok {
		return "", false
	}

	cachedValue := value.(CachedValue)

	if cachedValue.expiration.Before(time.Now()) {
		c.cache.Delete(breed)
		return "", false
	}

	return cachedValue.value, true
}

func (c *InMemoryCache) Set(breed, url string) {
	c.cache.Store(breed, CachedValue{value: url, expiration: time.Now().Add(c.duration)})
}
