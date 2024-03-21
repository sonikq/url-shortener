package storage

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

type memoryStorage struct {
	items map[string]Item
	mu    sync.RWMutex
}

type OptionsMemoryStorage func(m *memoryStorage)

func newMemoryStorage(opts ...OptionsMemoryStorage) *memoryStorage {
	c := &memoryStorage{
		items: make(map[string]Item),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithMemoryStorage(items map[string]Item) OptionsMemoryStorage {
	return func(m *memoryStorage) {
		m.items = items
	}
}

func (c *memoryStorage) Set(key string, value string, ttl time.Duration) {
	exp := time.Now().Add(ttl).UnixNano()

	c.mu.Lock()

	c.items[key] = Item{
		Object:     value,
		Expiration: exp,
	}

	c.mu.Unlock()
}

func (c *memoryStorage) Get(key string) (string, error) {
	c.mu.RLock()

	item, found := c.items[key]
	if !found {
		c.mu.RUnlock()
		return "", fmt.Errorf("access denied")
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return "", fmt.Errorf("memoryStorage time expired")
		}
	}
	c.mu.RUnlock()
	return item.Object, nil
}

func (c *memoryStorage) Flush() {
	c.mu.Lock()
	c.items = make(map[string]Item)
	c.mu.Unlock()
}
