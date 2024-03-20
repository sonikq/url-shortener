package repositories

import (
	"context"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/pkg/storage"
)

type IUserRepo interface {
	ShorteningLink(request user.ShorteningLinkRequest) user.ShorteningLinkResponse
	GetFullLinkByID(request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse
	ShorteningLinkJSON(request user.ShorteningLinkJSONRequest) user.ShorteningLinkJSONResponse
	PingDB(ctx context.Context) error
}

type Repository struct {
	IUserRepo
}

func NewRepository(storage *storage.Storage) *Repository {
	return &Repository{
		IUserRepo: NewUserRepo(storage),
	}
}
