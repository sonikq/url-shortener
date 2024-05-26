package repositories

import (
	"context"

	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/pkg/storage"
)

// IUserRepo -
type IUserRepo interface {
	ShorteningLink(ctx context.Context, request user.ShorteningLinkRequest) user.ShorteningLinkResponse
	GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse
	ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest) user.ShorteningLinkJSONResponse
	PingDB(ctx context.Context) error
	ShorteningBatchLinks(ctx context.Context, request user.ShorteningBatchLinksRequest) user.ShorteningBatchLinksResponse
	GetBatchByUserID(ctx context.Context, request user.GetBatchByUserIDRequest) user.GetBatchByUserIDResponse
}

// Repository -
type Repository struct {
	IUserRepo
}

// NewRepository -
func NewRepository(storage *storage.Storage) *Repository {
	return &Repository{
		IUserRepo: NewUserRepo(storage),
	}
}
