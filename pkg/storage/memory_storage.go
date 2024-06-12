package storage

import (
	"context"
	"fmt"
	"github.com/sonikq/url-shortener/internal/app/models"
	"sync"
	"time"
)

// Item -
type Item struct {
	Object     string
	UserID     string
	IsDeleted  bool
	Expiration int64
}

// Expired -
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

// OptionsMemoryStorage -
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

// WithMemoryStorage -
func WithMemoryStorage(items map[string]Item) OptionsMemoryStorage {
	return func(m *memoryStorage) {
		m.items = items
	}
}

// Set -
func (c *memoryStorage) Set(_ context.Context, data map[string]Item) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, value := range data {
		c.items[key] = value
	}

	return nil
}

// Get -
func (c *memoryStorage) Get(_ context.Context, alias string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[alias]
	if !found {
		return "", fmt.Errorf("access denied")
	}

	if item.IsDeleted {
		return "", models.ErrGetDeletedLink
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return "", fmt.Errorf("memoryStorage time expired")
		}
	}
	return item.Object, nil
}

// GetShortURL -
func (c *memoryStorage) GetShortURL(_ context.Context, originalURL string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for key, value := range c.items {
		if value.Object == originalURL {
			if value.Expiration > 0 {
				if time.Now().UnixNano() > value.Expiration {
					return "", fmt.Errorf("memoryStorage time expired")
				}
			}
			return key, nil
		}
	}

	return "", nil
}

// DeleteBatch -
func (c *memoryStorage) DeleteBatch(_ context.Context, urls []string, userID string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

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

	return nil
}

// GetBatchByUserID -
func (c *memoryStorage) GetBatchByUserID(_ context.Context, userID string) (map[string]Item, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	batch := make(map[string]Item)

	for key, item := range c.items {
		if item.UserID == userID {
			batch[key] = item
		}
	}

	return batch, nil
}

// Ping -
func (c *memoryStorage) Ping(_ context.Context) error {
	return fmt.Errorf("currently in use memory storage, not db")
}

// Close -
func (c *memoryStorage) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]Item)
}
