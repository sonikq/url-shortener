package main

import (
	"context"
	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/handlers"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/repositories"
	http2 "github.com/sonikq/url-shortener/internal/app/servers/http"
	"github.com/sonikq/url-shortener/internal/app/services"
	"github.com/sonikq/url-shortener/pkg/storage"
	lg "log"
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
		err := logger.CleanUp(log)
		log.Fatal("failed to cleanup logs", logger.Error(err))
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
		if err = server.Run(); err != nil {
			log.Fatal("failed to run http server")
		}
	}()

	lg.Println("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatal("failed to stop server")
	}
}

func initStorage(cfg cfg.Config) (*storage.Storage, error) {
	var storageOptions []storage.OptionsStorage
	if cfg.DatabaseDSN != "" {

		storageOptions = append(storageOptions, storage.WithDB(context.Background(), cfg.DatabaseDSN))
	}

	if cfg.FileStoragePath == "" {
		storageOptions = append(storageOptions, storage.RestoreFile(cfg.FileStoragePath))
		storageOptions = append(storageOptions, storage.WithFileStorage(cfg.FileStoragePath))
		return storage.NewStorage(storageOptions...)
	}

	return storage.NewStorage(storageOptions...)
}
