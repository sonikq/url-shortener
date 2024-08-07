package storage

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// FileStorage -
type FileStorage interface {
	SaveToFile(items map[string]Item) error
}

// IStorage -
type IStorage interface {
	Set(ctx context.Context, data map[string]Item) error
	Get(ctx context.Context, alias string) (string, error)
	GetShortURL(ctx context.Context, originalURL string) (string, error)
	Ping(ctx context.Context) error
	GetBatchByUserID(ctx context.Context, userID string) (map[string]Item, error)
	DeleteBatch(ctx context.Context, urls []string, userID string) error
	GetStats(ctx context.Context) (int64, int64, error)
	Close()
}

// Storage -
type Storage struct {
	File FileStorage
	IStorage
}

// OptionsStorage -
type OptionsStorage func(s *Storage) error

// NewStorage -
func NewStorage(opts ...OptionsStorage) (*Storage, error) {
	s := &Storage{}
	s.IStorage = newMemoryStorage()
	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// WithDB -
func WithDB(ctx context.Context, dsn string, dbPoolWorkers int) OptionsStorage {
	return func(s *Storage) error {
		var err error
		s.IStorage, err = newDB(ctx, dsn, dbPoolWorkers)
		return err
	}
}

// WithFileStorage -
func WithFileStorage(path string) OptionsStorage {
	return func(s *Storage) error {
		var err error
		s.File, err = newFileStorage(path)
		if err != nil {
			return err
		}
		return nil
	}
}

// RestoreFile -
func RestoreFile(ctx context.Context, filename string) OptionsStorage {
	return func(s *Storage) error {
		file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				fmt.Printf("cant close file: %s", err.Error())
			}
		}(file)
		if err != nil {
			return fmt.Errorf("cant open file: %s", err.Error())
		}

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		data := scanner.Bytes()

		if data == nil {
			return nil
		}

		itemsMap := make(map[string]Item)

		err = json.Unmarshal(data, &itemsMap)
		if err != nil {
			return fmt.Errorf("cant unmarshal objects from file: %s", err.Error())
		}

		if s.IStorage != nil {
			err = s.Set(ctx, itemsMap)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
