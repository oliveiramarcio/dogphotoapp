package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryCache_Get_ReturnsValueFromCache(t *testing.T) {
	cache := &InMemoryCache{
		cache:    sync.Map{},
		duration: time.Minute,
	}
	breed := "dog"
	url := "https://example.com/dog.jpg"
	cache.Set(breed, url)

	value, ok := cache.Get(breed)

	assert.True(t, ok)
	assert.Equal(t, url, value)
}

func TestInMemoryCache_Get_ReturnsEmptyStringWhenValueNotPresent(t *testing.T) {
	cache := &InMemoryCache{
		cache:    sync.Map{},
		duration: time.Minute,
	}
	breed := "cat"

	value, ok := cache.Get(breed)

	assert.False(t, ok)
	assert.Empty(t, value)
}

func TestInMemoryCache_Get_ReturnsEmptyStringWhenValueExpired(t *testing.T) {
	cache := &InMemoryCache{
		cache:    sync.Map{},
		duration: -time.Minute, // Negative duration to make the value immediately expired
	}
	breed := "dog"
	url := "https://example.com/dog.jpg"
	cache.Set(breed, url)

	value, ok := cache.Get(breed)

	assert.False(t, ok)
	assert.Empty(t, value)
}

func TestInMemoryCache_Get_DeletesExpiredValueFromCache(t *testing.T) {
	cache := &InMemoryCache{
		cache:    sync.Map{},
		duration: -time.Minute, // Negative duration to make the value immediately expired
	}
	breed := "dog"
	url := "https://example.com/dog.jpg"
	cache.Set(breed, url)

	cache.Get(breed)

	_, ok := cache.cache.Load(breed)
	assert.False(t, ok, "value was not deleted from cache")
}

func TestInMemoryCache_Set_SetsValueInCache(t *testing.T) {
	cache := &InMemoryCache{
		cache:    sync.Map{},
		duration: time.Minute,
	}
	breed := "dog"
	url := "https://example.com/dog.jpg"

	cache.Set(breed, url)

	value, ok := cache.cache.Load(breed)
	assert.True(t, ok, "value not found in cache")
	cachedValue := value.(CachedValue)
	assert.Equal(t, url, cachedValue.value, "incorrect value in cache")
}

func TestInMemoryCache_Set_UpdatesExistingValueInCache(t *testing.T) {
	cache := &InMemoryCache{
		cache:    sync.Map{},
		duration: time.Minute,
	}
	breed := "dog"
	oldURL := "https://example.com/dog.jpg"
	newURL := "https://example.com/dog-new.jpg"
	cache.Set(breed, oldURL)

	cache.Set(breed, newURL)

	value, ok := cache.cache.Load(breed)
	assert.True(t, ok, "value not found in cache")
	cachedValue := value.(CachedValue)
	assert.Equal(t, newURL, cachedValue.value, "incorrect value in cache")
}
