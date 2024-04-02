package utils

import (
	"github.com/sonikq/url-shortener/pkg/storage"
	"time"
)

func ConvertDataToStore(alias, originalURL, userID string) map[string]storage.Item {
	mapToStore := make(map[string]storage.Item)
	itemToStoreInDB := storage.Item{
		Object:     originalURL,
		Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
		UserID:     userID,
	}
	mapToStore[alias] = itemToStoreInDB
	return mapToStore
}
