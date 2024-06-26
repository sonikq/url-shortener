package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type fileStorage struct {
	file *os.File
}

func newFileStorage(path string) (*fileStorage, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}

	return &fileStorage{
		file: file,
	}, nil
}

// SaveToFile -
func (f *fileStorage) SaveToFile(items map[string]Item) error {
	for key, item := range items {
		intermediateMap := make(map[string]Item)
		intermediateMap[key] = item
		data, err := json.Marshal(intermediateMap)
		if err != nil {
			return err
		}
		data = append(data, '\n')

		_, err = f.file.Write(data)
		if err != nil {
			return fmt.Errorf("error in saving file: %s", err.Error())
		}
	}
	return nil
}
