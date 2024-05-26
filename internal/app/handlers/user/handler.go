package user

import (
	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/services"
	"github.com/sonikq/url-shortener/internal/app/workers"
)

// HandlerConfig -
type HandlerConfig struct {
	Conf    cfg.Config
	Logger  logger.Logger
	Service *services.Service
	Worker  *workers.Worker
}

// Handler -
type Handler struct {
	config  cfg.Config
	log     logger.Logger
	service *services.Service
	worker  *workers.Worker
}

// New -
func New(cfg *HandlerConfig) *Handler {
	return &Handler{
		config:  cfg.Conf,
		log:     cfg.Logger,
		service: cfg.Service,
		worker:  cfg.Worker,
	}
}
