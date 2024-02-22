package main

import (
	"context"
	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/handlers"
	"github.com/sonikq/url-shortener/internal/app/repositories"
	http2 "github.com/sonikq/url-shortener/internal/app/servers/http"
	"github.com/sonikq/url-shortener/internal/app/services"
	"github.com/sonikq/url-shortener/pkg/cache"
	"log"
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
		log.Fatal("failed to initialize configuration")
	}

	cfg.ParseConfig(&config)

	_cache := cache.New()
	defer _cache.FlushCache()

	repo := repositories.NewRepository(_cache)

	service := services.NewService(repo)

	router := handlers.NewRouter(handlers.Option{
		Conf:    config,
		Cache:   _cache,
		Service: service,
	})

	server := http2.NewServer(config.HTTP.Port, router)

	go func() {
		if err = server.Run(); err != nil {
			log.Fatal("failed to run http server")
		}
	}()

	log.Println("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatal("failed to stop server")
	}
}
