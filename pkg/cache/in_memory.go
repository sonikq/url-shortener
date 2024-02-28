package cache

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
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

	if c.filePath != "" {
		err := c.SaveToFile()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error saving data to file: %v\n", err)
		}
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

func (c *Cache) RestoreFromFile() error {
	dir, _ := path.Split(c.filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0666)
		if err != nil {
			return fmt.Errorf("cant create directory: %s", err.Error())
		}
	}
	file, err := os.OpenFile(c.filePath, os.O_RDONLY|os.O_CREATE, 0666)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Errorf("cant close file: %s", err.Error())
		}
	}(file)
	if err != nil {
		return fmt.Errorf("cant open file: %s", err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := scanner.Bytes()

	err = json.Unmarshal(data, &c.items)
	if err != nil {
		return fmt.Errorf("cant unmarshal objects from file: %s", err.Error())
	}
	return nil
}

func (c *Cache) SaveToFile() error {
	file, err := os.OpenFile(c.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Errorf("cant close file: %s", err.Error())
		}
	}(file)
	if err != nil {
		return fmt.Errorf("cant open file: %s", err.Error())
	}

	data, _ := json.Marshal(c.items)

	data = append(data, '\n')

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error in saving file: %s", err.Error())
	}
	return nil
}

func (c *Cache) FlushCache() {
	c.mu.Lock()
	c.items = make(map[string]Item)
	c.mu.Unlock()
}
