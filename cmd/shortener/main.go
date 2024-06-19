package main

import (
	"context"
	"errors"
	"fmt"
	lg "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/handlers"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/repositories"
	http2 "github.com/sonikq/url-shortener/internal/app/servers/http"
	"github.com/sonikq/url-shortener/internal/app/services"
	"github.com/sonikq/url-shortener/internal/app/workers"
	"github.com/sonikq/url-shortener/pkg/storage"
)

var buildVersion string = "N/A"
var buildDate string = "N/A"
var buildCommit string = "N/A"

// printBuildInfo prints the build information.
func printBuildInfo() {
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
}

func main() {
	printBuildInfo()

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
	//defer func() {
	//	err = logger.CleanUp(log)
	//	log.Info("failed to cleanup logs", logger.Error(err))
	//}()

	store, err := initStorage(config)
	if err != nil {
		log.Fatal("failed to initialize storage", logger.Error(err))
	}

	repo := repositories.NewRepository(store)

	service := services.NewService(repo)

	pool := make(chan workers.Pool)
	defer close(pool)

	worker := workers.NewWorker(pool, store)
	go worker.Run()

	router := handlers.NewRouter(handlers.Option{
		Conf:    config,
		Cache:   store,
		Logger:  log,
		Service: service,
		Worker:  worker,
	})

	server := http2.NewServer(config.HTTP, router)

	go func() {
		err = server.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("failed to run http server")
		}
	}()

	lg.Println("Server started...")

	idleConnsClosed := make(chan struct{})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSEGV)

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		<-quit
		if err = server.Shutdown(ctxShutdown); err != nil {
			log.Error("error in shutting down server")
		}
		close(idleConnsClosed)
	}()

	<-idleConnsClosed

	log.Info("server shutdown gracefully")
}

func initStorage(cfg cfg.Config) (*storage.Storage, error) {
	var storageOptions []storage.OptionsStorage
	if cfg.DatabaseDSN != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		storageOptions = append(storageOptions, storage.WithDB(ctx, cfg.DatabaseDSN, cfg.DBPoolWorkers))
	}

	if cfg.FileStoragePath != "" {
		storageOptions = append(storageOptions, storage.RestoreFile(context.Background(), cfg.FileStoragePath))
		storageOptions = append(storageOptions, storage.WithFileStorage(cfg.FileStoragePath))
		return storage.NewStorage(storageOptions...)
	}

	return storage.NewStorage(storageOptions...)
}
