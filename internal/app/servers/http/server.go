package http

import (
	"context"
	"github.com/sonikq/url-shortener/configs/app"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"time"
)

// Server -
type Server struct {
	httpServer *http.Server
}

// NewServer -
func NewServer(config app.HTTPConfig, handler http.Handler) *Server {
	// HTTPS
	if config.EnableHTTPS != "" {
		manager := &autocert.Manager{
			Cache:      autocert.DirCache("cache-dir"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(config.ServerAddress),
		}
		return &Server{
			httpServer: &http.Server{
				Addr:           ":" + config.Port,
				Handler:        handler,
				ReadTimeout:    15 * time.Minute,
				WriteTimeout:   15 * time.Minute,
				MaxHeaderBytes: 1 << 20,
				TLSConfig:      manager.TLSConfig(),
			},
		}
	}

	// HTTP
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + config.Port,
			Handler:        handler,
			ReadTimeout:    15 * time.Minute,
			WriteTimeout:   15 * time.Minute,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

// Run -
func (s *Server) Run() error {
	if s.httpServer.TLSConfig != nil {
		return s.httpServer.ListenAndServeTLS("", "")
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown -
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
