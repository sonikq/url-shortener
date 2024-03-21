package storage

import (
	"context"
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

func (c *memoryStorage) Set(ctx context.Context, data map[string]Item) (*string, error) {
	c.mu.Lock()
	for key, value := range data {
		c.items[key] = value
	}

	c.mu.Unlock()
	return nil, nil
}

func (c *memoryStorage) Get(ctx context.Context, alias string) (string, error) {
	c.mu.RLock()

	item, found := c.items[alias]
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

func (c *memoryStorage) Ping(ctx context.Context) error {
	return fmt.Errorf("currently in use memory storage, not db")
}

func (c *memoryStorage) Close() {
	c.mu.Lock()
	c.items = make(map[string]Item)
	c.mu.Unlock()
}
