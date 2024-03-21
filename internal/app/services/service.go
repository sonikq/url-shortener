package services

import (
	"context"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/repositories"
)

type IUserService interface {
	ShorteningLink(ctx context.Context, request user.ShorteningLinkRequest, response chan user.ShorteningLinkResponse)
	GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest, response chan user.GetFullLinkByIDResponse)
	ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest, response chan user.ShorteningLinkJSONResponse)
	PingDB(ctx context.Context) error
}

type Service struct {
	IUserService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IUserService: NewUserService(repos.IUserRepo),
	}
}
