package repositories

import (
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/pkg/cache"
)

type IUserRepo interface {
	ShorteningLink(request user.ShorteningLinkRequest) user.ShorteningLinkResponse
	GetFullLinkByID(request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse
}

type Repository struct {
	IUserRepo
}

func NewRepository(c *cache.Cache) *Repository {
	return &Repository{
		IUserRepo: NewUserRepo(c),
	}
}
