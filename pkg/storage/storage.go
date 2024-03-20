package storage

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type MemoryStorage interface {
	Set(key string, value string, ttl time.Duration)
	Get(key string) (string, error)
	Flush()
}

type FileStorage interface {
	SaveToFile(items map[string]Item) error
}

type DB interface {
	Ping(ctx context.Context) error
	Close()
}

type Storage struct {
	File   FileStorage
	Memory MemoryStorage
	DB     DB
}

type OptionsStorage func(s *Storage) error

func NewStorage(opts ...OptionsStorage) (*Storage, error) {
	s := &Storage{
		Memory: newMemoryStorage(),
	}
	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func WithDB(ctx context.Context, dsn string) OptionsStorage {
	return func(s *Storage) error {
		var err error
		s.DB, err = newPostgresClient(ctx, dsn)
		return err
	}
}

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

func RestoreFile(filename string) OptionsStorage {
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

		var options []OptionsMemoryStorage

		options = append(options, WithMemoryStorage(itemsMap))

		s.Memory = newMemoryStorage(options...)

		return nil
	}
}
