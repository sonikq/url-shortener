package user

import (
	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/services"
)

type HandlerConfig struct {
	Conf    cfg.Config
	Service *services.Service
}

type Handler struct {
	config  cfg.Config
	service *services.Service
}

func New(cfg *HandlerConfig) *Handler {
	return &Handler{
		config:  cfg.Conf,
		service: cfg.Service,
	}
}
