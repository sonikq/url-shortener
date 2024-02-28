package cache

import (
	"encoding/json"
	"fmt"
	"os"
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
	items    map[string]Item
	mu       sync.RWMutex
	filePath string
}

func New(filePath string) *Cache {
	return &Cache{
		newCache(make(map[string]Item), filePath),
	}
}

func newCache(m map[string]Item, pathToFile string) *cache {
	c := &cache{
		items:    m,
		filePath: pathToFile,
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

	// Сохранение данных в файл
	err := c.SaveToFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving data to file: %v\n", err)
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

func (c *Cache) LoadFromFile() error {
	if c.filePath == "" {
		return nil
	}

	if _, err := os.Stat(c.filePath); os.IsNotExist(err) {
		file, err := os.Create(c.filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	file, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		return nil
	}

	if err = json.NewDecoder(file).Decode(&c.items); err != nil {
		return err
	}

	return nil
}

func (c *Cache) SaveToFile() error {
	var file *os.File

	if c.filePath == "" {
		return nil
	}
	if _, err := os.Stat(c.filePath); os.IsNotExist(err) {
		file, err = os.Create(c.filePath)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	err := json.NewEncoder(file).Encode(c.items)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) FlushCache() {
	c.mu.Lock()
	c.items = make(map[string]Item)
	c.mu.Unlock()
}
