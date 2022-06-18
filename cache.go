package study

import (
	"sync"
	"time"
)

type CacheLoader interface {
	load(key string) (any, error)
}

type loadFn func(key string) (any, error)

func (f loadFn) load(key string) (any, error) {
	return f(key)
}

type Cache struct {
	sync.RWMutex
	ttl         time.Duration
	entities    map[string]Entity
	CacheLoader CacheLoader
}

type Entity struct {
	value any
	ttl   int64
}

func NewCache(ttl time.Duration, loader CacheLoader) *Cache {
	return &Cache{
		ttl:         ttl,
		entities:    make(map[string]Entity),
		CacheLoader: loader,
	}
}

func (c *Cache) Set(key string, value any) {
	c.Lock()
	defer c.Unlock()
	c.entities[key] = Entity{value: value, ttl: time.Now().Add(c.ttl).UnixNano()}
}

func (c *Cache) Get(key string) (any, error) {
	c.RLock()
	defer c.RUnlock()
	_, ok := c.entities[key]
	if !ok {
		value, err := c.CacheLoader.load(key)
		if err != nil {
			c.Set(key, value)
		} else {
			return nil, err
		}
	}
	return c.entities[key], nil
}
