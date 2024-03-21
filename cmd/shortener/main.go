package main

import (
	"context"
	"errors"
	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/handlers"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/repositories"
	http2 "github.com/sonikq/url-shortener/internal/app/servers/http"
	"github.com/sonikq/url-shortener/internal/app/services"
	"github.com/sonikq/url-shortener/pkg/storage"
	lg "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		config cfg.Config // Configurations
		err    error
	)

	config, err = cfg.Load("configs/app/.env")
	if err != nil {
		lg.Fatal("failed to initialize configuration")
	}

	cfg.ParseConfig(&config)

	// Logger
	log := logger.New(config.LogLevel, config.ServiceName)
	defer func() {
		err = logger.CleanUp(log)
		log.Info("failed to cleanup logs", logger.Error(err))
	}()

	store, err := initStorage(config)
	defer store.Memory.Flush()
	if err != nil {
		log.Fatal("failed to initialize storage", logger.Error(err))
	}

	repo := repositories.NewRepository(store)

	service := services.NewService(repo)

	router := handlers.NewRouter(handlers.Option{
		Conf:    config,
		Cache:   store,
		Logger:  log,
		Service: service,
	})

	server := http2.NewServer(config.HTTP.Port, router)

	go func() {
		err = server.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("failed to run http server")
		}
	}()

	lg.Println("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(ctxShutdown); err != nil {
		log.Error("error in shutting down server")
	} else {
		log.Info("server stopped successfully")
	}
}

func initStorage(cfg cfg.Config) (*storage.Storage, error) {
	var storageOptions []storage.OptionsStorage
	if cfg.DatabaseDSN != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()
		storageOptions = append(storageOptions, storage.WithDB(ctx, cfg.DatabaseDSN))
	}

	if cfg.FileStoragePath != "" {
		storageOptions = append(storageOptions, storage.RestoreFile(cfg.FileStoragePath))
		storageOptions = append(storageOptions, storage.WithFileStorage(cfg.FileStoragePath))
		return storage.NewStorage(storageOptions...)
	}

	return storage.NewStorage(storageOptions...)
}
