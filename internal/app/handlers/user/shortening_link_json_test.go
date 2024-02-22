package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/services"
	"testing"
)

func TestHandler_ShorteningLinkJSON(t *testing.T) {
	type fields struct {
		config  app.Config
		log     logger.Logger
		service *services.Service
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				config:  tt.fields.config,
				log:     tt.fields.log,
				service: tt.fields.service,
			}
			h.ShorteningLinkJSON(tt.args.ctx)
		})
	}
}
