package services

import (
	"context"

	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/repositories"
)

type UserService struct {
	repo repositories.IUserRepo
}

func NewUserService(repo repositories.IUserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) ShorteningLink(ctx context.Context, request user.ShorteningLinkRequest) user.ShorteningLinkResponse {
	return s.repo.ShorteningLink(ctx, request)
}

func (s *UserService) GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse {
	return s.repo.GetFullLinkByID(ctx, request)

}

func (s *UserService) ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest) user.ShorteningLinkJSONResponse {
	return s.repo.ShorteningLinkJSON(ctx, request)
}

func (s *UserService) ShorteningBatchLinks(ctx context.Context, request user.ShorteningBatchLinksRequest) user.ShorteningBatchLinksResponse {
	return s.repo.ShorteningBatchLinks(ctx, request)
}

func (s *UserService) GetBatchByUserID(ctx context.Context, request user.GetBatchByUserIDRequest) user.GetBatchByUserIDResponse {
	return s.repo.GetBatchByUserID(ctx, request)
}

func (s *UserService) PingDB(ctx context.Context) error {
	return s.repo.PingDB(ctx)
}
