package services

import (
	"context"

	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/repositories"
)

// IUserService -
type IUserService interface {
	ShorteningLink(ctx context.Context, request user.ShorteningLinkRequest) user.ShorteningLinkResponse
	GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse
	ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest) user.ShorteningLinkJSONResponse
	PingDB(ctx context.Context) error
	ShorteningBatchLinks(ctx context.Context, request user.ShorteningBatchLinksRequest) user.ShorteningBatchLinksResponse
	GetBatchByUserID(ctx context.Context, request user.GetBatchByUserIDRequest) user.GetBatchByUserIDResponse
	GetStats(ctx context.Context) user.GetStatsResponse
}

// Service -
type Service struct {
	IUserService
}

// NewService -
func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IUserService: NewUserService(repos.IUserRepo),
	}
}
