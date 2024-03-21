package repositories

import (
	"context"
	"fmt"
	"github.com/sonikq/url-shortener/internal/app/models"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/utils"
	"github.com/sonikq/url-shortener/pkg/storage"
	"time"
)

type UserRepo struct {
	storage *storage.Storage
}

func NewUserRepo(storage *storage.Storage) *UserRepo {
	return &UserRepo{
		storage: storage,
	}
}

func (r *UserRepo) ShorteningLink(ctx context.Context, request user.ShorteningLinkRequest) user.ShorteningLinkResponse {
	alias := utils.RandomString(sizeOfAlias)

	result := httpPrefix + request.RequestURL + alias
	if r.storage.DB != nil {
		itemToStoreInDB := storage.Item{
			Object:     request.ShorteningLink,
			Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
		}
		mapToStoreInDB := make(map[string]storage.Item)
		mapToStoreInDB[alias] = itemToStoreInDB
		err := r.storage.DB.Set(ctx, mapToStoreInDB)
		if err != nil {
			return user.ShorteningLinkResponse{
				Code:   500,
				Status: fail,
				Error: &models.Err{
					Source:  "db_storage",
					Message: err.Error(),
				},
				Response: nil,
			}
		}
	} else {
		r.storage.Memory.Set(alias, request.ShorteningLink, 10*time.Minute)
	}

	if r.storage.File != nil {
		itemToStoreInFile := storage.Item{
			Object:     request.ShorteningLink,
			Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
		}
		mapToStoreInFile := make(map[string]storage.Item)
		mapToStoreInFile[alias] = itemToStoreInFile
		err := r.storage.File.SaveToFile(mapToStoreInFile)
		if err != nil {
			return user.ShorteningLinkResponse{
				Code:   500,
				Status: fail,
				Error: &models.Err{
					Source:  "file_storage",
					Message: err.Error(),
				},
				Response: nil,
			}
		}
	}

	return user.ShorteningLinkResponse{
		Code:     201,
		Status:   success,
		Error:    nil,
		Response: &result,
	}
}

func (r *UserRepo) ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest) user.ShorteningLinkJSONResponse {
	alias := utils.RandomString(sizeOfAlias)

	result := request.BaseURL + "/" + alias

	if r.storage.DB != nil {
		itemToStoreInDB := storage.Item{
			Object:     request.ShorteningLink.URL,
			Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
		}
		mapToStoreInDB := make(map[string]storage.Item)
		mapToStoreInDB[alias] = itemToStoreInDB
		err := r.storage.DB.Set(ctx, mapToStoreInDB)
		if err != nil {
			return user.ShorteningLinkJSONResponse{
				Code:   500,
				Status: fail,
				Error: &models.Err{
					Source:  "db_storage",
					Message: err.Error(),
				},
				Response: user.ShortenLinkJSONResponseBody{},
			}
		}
	} else {
		r.storage.Memory.Set(alias, request.ShorteningLink.URL, 10*time.Minute)
	}

	if r.storage.File != nil {
		itemToStoreInFile := storage.Item{
			Object:     request.ShorteningLink.URL,
			Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
		}
		mapToStoreInFile := make(map[string]storage.Item)
		mapToStoreInFile[alias] = itemToStoreInFile
		err := r.storage.File.SaveToFile(mapToStoreInFile)
		if err != nil {
			return user.ShorteningLinkJSONResponse{
				Code:   500,
				Status: fail,
				Error: &models.Err{
					Source:  "file_storage",
					Message: err.Error(),
				},
				Response: user.ShortenLinkJSONResponseBody{},
			}
		}
	}

	return user.ShorteningLinkJSONResponse{
		Code:     201,
		Status:   success,
		Error:    nil,
		Response: user.ShortenLinkJSONResponseBody{Result: result},
	}
}

func (r *UserRepo) GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse {
	var fullLink string
	var err error
	if r.storage.DB != nil {
		fullLink, err = r.storage.DB.Get(ctx, request.ShortLinkID)
		if err != nil {
			return user.GetFullLinkByIDResponse{
				Code:   500,
				Status: fail,
				Error: &models.Err{
					Source:  "db_storage",
					Message: err.Error(),
				},
				Response: nil,
			}
		}
	} else {
		fullLink, err = r.storage.Memory.Get(request.ShortLinkID)
		if err != nil {
			return user.GetFullLinkByIDResponse{
				Code:   500,
				Status: fail,
				Error: &models.Err{
					Source:  "memory_storage",
					Message: err.Error(),
				},
				Response: nil,
			}
		}
	}

	return user.GetFullLinkByIDResponse{
		Code:     307,
		Status:   success,
		Error:    nil,
		Response: &fullLink,
	}
}

func (r *UserRepo) PingDB(ctx context.Context) error {
	if r.storage.DB != nil {
		fmt.Printf("error from db-ping: %v", r.storage.DB.Ping(ctx))
		return r.storage.DB.Ping(ctx)
	}
	return fmt.Errorf("the database is not responding")
}
