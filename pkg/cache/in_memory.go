package cache

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	Object     string
	Expiration int64
}

func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	*cache
}

type cache struct {
	items map[string]Item
	mu    sync.RWMutex
}

func New() *Cache {
	return &Cache{
		newCache(make(map[string]Item)),
	}
}

func newCache(m map[string]Item) *cache {
	c := &cache{
		items: m,
	}
	return c
}

func (c *Cache) Set(key string, value string, ttl time.Duration) {
	exp := time.Now().Add(ttl).UnixNano()

	c.mu.Lock()

	c.items[key] = Item{
		Object:     value,
		Expiration: exp,
	}

	c.mu.Unlock()
}

func (c *Cache) Get(key string) (string, error) {
	c.mu.RLock()

	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return "", fmt.Errorf("access denied")
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return "", fmt.Errorf("cache time expired")
		}
	}
	c.mu.RUnlock()
	return item.Object, nil
}

func (c *Cache) FlushCache() {
	c.mu.Lock()
	c.items = make(map[string]Item)
	c.mu.Unlock()
}
