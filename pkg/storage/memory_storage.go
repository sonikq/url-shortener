package storage

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Item struct {
	Object     string
	UserID     string
	IsDeleted  bool
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

func (c *memoryStorage) Set(ctx context.Context, data map[string]Item) error {
	c.mu.Lock()
	for key, value := range data {
		c.items[key] = value
	}

	c.mu.Unlock()
	return nil
}

func (c *memoryStorage) Get(ctx context.Context, alias string) (string, error) {
	c.mu.RLock()

	item, found := c.items[alias]
	if !found {
		c.mu.RUnlock()
		return "", fmt.Errorf("access denied")
	}

	if item.IsDeleted {
		c.mu.RUnlock()
		return "", ErrGetDeletedLink
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

func (c *memoryStorage) GetShortURL(ctx context.Context, originalURL string) (string, error) {
	c.mu.RLock()
	for key, value := range c.items {
		if value.Object == originalURL {
			if value.Expiration > 0 {
				if time.Now().UnixNano() > value.Expiration {
					c.mu.RUnlock()
					return "", fmt.Errorf("memoryStorage time expired")
				}
			}
			return key, nil
		}
	}
	c.mu.RUnlock()
	return "", nil
}

func (c *memoryStorage) DeleteBatch(ctx context.Context, urls []string, userID string) error {
	c.mu.Lock()
	for _, value := range urls {
		if c.items[value].UserID == userID {
			c.items[value] = Item{
				Object:     c.items[value].Object,
				Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
				UserID:     userID,
				IsDeleted:  true,
			}
		}
	}

	c.mu.Unlock()

	return nil
}

func (c *memoryStorage) GetBatchByUserID(ctx context.Context, userID string) (map[string]Item, error) {
	c.mu.Lock()
	batch := make(map[string]Item)

	for key, item := range c.items {
		if item.UserID == userID {
			batch[key] = item
		}
	}
	c.mu.Unlock()
	return batch, nil
}

func (c *memoryStorage) Ping(ctx context.Context) error {
	return fmt.Errorf("currently in use memory storage, not db")
}

func (c *memoryStorage) Close() {
	c.mu.Lock()
	c.items = make(map[string]Item)
	c.mu.Unlock()
}
