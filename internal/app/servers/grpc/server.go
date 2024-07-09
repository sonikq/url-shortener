package grpc

import (
	"github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/pkg/interceptors"
	pb "github.com/sonikq/url-shortener/internal/app/proto"
	"github.com/sonikq/url-shortener/internal/app/repositories"
	"github.com/sonikq/url-shortener/internal/app/services"
	"google.golang.org/grpc"
	"net"
)

// Server -
type Server struct {
	grpcServer *grpc.Server
}

// NewServer -
func NewServer(conf app.Config, repo repositories.IUserRepo) *Server {
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors.UnaryServerInterceptorOpts(conf.TrustedSubnet)))

	pb.RegisterShortenerServer(server, &services.ServiceGrpc{Repo: repo, BaseURL: conf.BaseURL})
	return &Server{server}
}

// Run -
func (s *Server) Run(port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	if err = s.grpcServer.Serve(listen); err != nil {
		return err
	}
	return nil
}

// Shutdown -
func (s *Server) Shutdown() {
	s.grpcServer.GracefulStop()
}
