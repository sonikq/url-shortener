package workers

import (
	"context"

	"github.com/sonikq/url-shortener/pkg/storage"
)

type Worker struct {
	pool  chan Pool
	store *storage.Storage
}

type Pool struct {
	urls   []string
	err    chan error
	userID string
}

func NewWorker(urlsChan chan Pool, store *storage.Storage) *Worker {
	return &Worker{
		pool:  urlsChan,
		store: store,
	}
}

func (w *Worker) DeleteURLs(urls []string, userID string) error {
	errChan := make(chan error)

	// Создаем Pool и отправляем его в канал
	w.pool <- Pool{
		urls:   urls,
		err:    errChan,
		userID: userID,
	}
	err := <-errChan
	close(errChan)
	return err
}

func (w *Worker) Run() {
	for p := range w.pool {
		err := w.store.DeleteBatch(context.Background(), p.urls, p.userID)
		p.err <- err
	}
}
