package services

import (
	"context"

	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/repositories"
)

// UserService -
type UserService struct {
	repo repositories.IUserRepo
}

// NewUserService -
func NewUserService(repo repositories.IUserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

// ShorteningLink -
func (s *UserService) ShorteningLink(ctx context.Context, request user.ShorteningLinkRequest) user.ShorteningLinkResponse {
	return s.repo.ShorteningLink(ctx, request)
}

// GetFullLinkByID -
func (s *UserService) GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse {
	return s.repo.GetFullLinkByID(ctx, request)

}

// ShorteningLinkJSON -
func (s *UserService) ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest) user.ShorteningLinkJSONResponse {
	return s.repo.ShorteningLinkJSON(ctx, request)
}

// ShorteningBatchLinks -
func (s *UserService) ShorteningBatchLinks(ctx context.Context, request user.ShorteningBatchLinksRequest) user.ShorteningBatchLinksResponse {
	return s.repo.ShorteningBatchLinks(ctx, request)
}

// GetBatchByUserID -
func (s *UserService) GetBatchByUserID(ctx context.Context, request user.GetBatchByUserIDRequest) user.GetBatchByUserIDResponse {
	return s.repo.GetBatchByUserID(ctx, request)
}

// GetStats -
func (s *UserService) GetStats(ctx context.Context) user.GetStatsResponse {
	return s.repo.GetStats(ctx)
}

// PingDB -
func (s *UserService) PingDB(ctx context.Context) error {
	return s.repo.PingDB(ctx)
}
