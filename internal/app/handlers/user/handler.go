package user

import (
	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/services"
)

type HandlerConfig struct {
	Conf    cfg.Config
	Logger  logger.Logger
	Service *services.Service
}

type Handler struct {
	config  cfg.Config
	log     logger.Logger
	service *services.Service
}

func New(cfg *HandlerConfig) *Handler {
	return &Handler{
		config:  cfg.Conf,
		log:     cfg.Logger,
		service: cfg.Service,
	}
}
