package http

import (
	"context"
	"net/http"
	"time"
)

// Server -
type Server struct {
	httpServer *http.Server
}

// NewServer -
func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			ReadTimeout:    15 * time.Minute,
			WriteTimeout:   15 * time.Minute,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

// Run -
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// Shutdown -
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
