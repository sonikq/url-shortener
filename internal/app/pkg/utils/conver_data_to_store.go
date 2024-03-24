package utils

import (
	"github.com/sonikq/url-shortener/pkg/storage"
	"time"
)

func ConvertDataToStore(alias, originalURL string) map[string]storage.Item {
	mapToStore := make(map[string]storage.Item)
	itemToStoreInDB := storage.Item{
		Object:     originalURL,
		Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
	}
	mapToStore[alias] = itemToStoreInDB
	return mapToStore
}
