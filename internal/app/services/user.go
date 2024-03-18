package services

import (
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

func (s *UserService) ShorteningLink(request user.ShorteningLinkRequest, response chan user.ShorteningLinkResponse) {
	result := s.repo.ShorteningLink(request)

	response <- user.ShorteningLinkResponse{
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
		Response: result.Response,
	}
}

func (s *UserService) GetFullLinkByID(request user.GetFullLinkByIDRequest, response chan user.GetFullLinkByIDResponse) {
	result := s.repo.GetFullLinkByID(request)

	response <- user.GetFullLinkByIDResponse{
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
		Response: result.Response,
	}
}

func (s *UserService) ShorteningLinkJSON(request user.ShorteningLinkJSONRequest, response chan user.ShorteningLinkJSONResponse) {
	result := s.repo.ShorteningLinkJSON(request)

	response <- user.ShorteningLinkJSONResponse{
		Code:     result.Code,
		Status:   result.Status,
		Error:    result.Error,
		Response: result.Response,
	}
}
