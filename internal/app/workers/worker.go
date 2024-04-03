package workers

import (
	"context"
	"github.com/sonikq/url-shortener/pkg/storage"
)

type Worker struct {
	urlsChan chan deleteRequest
	store    *storage.Storage
}

type deleteRequest struct {
	urls   []string
	userID string
}

func NewWorker(store *storage.Storage) *Worker {
	// Инициализация канала
	urlsChan := make(chan deleteRequest)

	// Запуск горутины для обработки запросов на удаление
	go processDeleteRequests(urlsChan, store)

	return &Worker{
		urlsChan: urlsChan,
		store:    store,
	}
}

func (w *Worker) DeleteURLs(urls []string, userID string) {
	// Создаем deleteRequest и отправляем его в канал
	req := deleteRequest{
		urls:   urls,
		userID: userID,
	}
	w.urlsChan <- req
}

func processDeleteRequests(urlsChan <-chan deleteRequest, store *storage.Storage) {
	for req := range urlsChan {
		store.DeleteBatch(context.Background(), req.urls, req.userID)
	}
}
