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

func (s *UserService) ShorteningLink(ctx context.Context, request user.ShorteningLinkRequest, response chan user.ShorteningLinkResponse) {
	result := s.repo.ShorteningLink(ctx, request)

	response <- user.ShorteningLinkResponse{
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
		Response: result.Response,
	}
}

func (s *UserService) GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest, response chan user.GetFullLinkByIDResponse) {
	result := s.repo.GetFullLinkByID(ctx, request)

	response <- user.GetFullLinkByIDResponse{
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
		Response: result.Response,
	}
}

func (s *UserService) ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest, response chan user.ShorteningLinkJSONResponse) {
	result := s.repo.ShorteningLinkJSON(ctx, request)

	response <- user.ShorteningLinkJSONResponse{
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
		Response: result.Response,
	}
}

func (s *UserService) ShorteningBatchLinks(ctx context.Context, request user.ShorteningBatchLinksRequest, response chan user.ShorteningBatchLinksResponse) {
	result := s.repo.ShorteningBatchLinks(ctx, request)

	response <- user.ShorteningBatchLinksResponse{
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
		Response: result.Response,
	}
}

func (s *UserService) PingDB(ctx context.Context) error {
	return s.repo.PingDB(ctx)
}
