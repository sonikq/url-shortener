package repositories

import (
	"context"
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
	result := request.BaseURL + "/" + alias

	itemToStoreInDB := storage.Item{
		Object:     request.ShorteningLink,
		Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
	}
	mapToStoreInDB := make(map[string]storage.Item)
	mapToStoreInDB[alias] = itemToStoreInDB
	err := r.storage.Set(ctx, mapToStoreInDB)
	if err != nil {
		return user.ShorteningLinkResponse{
			Code:   500,
			Status: fail,
			Error: &models.Err{
				Source:  "storage",
				Message: err.Error(),
			},
			Response: nil,
		}
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

	itemToStoreInDB := storage.Item{
		Object:     request.ShorteningLink.URL,
		Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
	}
	mapToStoreInDB := make(map[string]storage.Item)
	mapToStoreInDB[alias] = itemToStoreInDB
	err := r.storage.Set(ctx, mapToStoreInDB)
	if err != nil {
		return user.ShorteningLinkJSONResponse{
			Code:   500,
			Status: fail,
			Error: &models.Err{
				Source:  "storage",
				Message: err.Error(),
			},
			Response: user.ShortenLinkJSONResponseBody{},
		}
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
	fullLink, err := r.storage.Get(ctx, request.ShortLinkID)
	if err != nil {
		return user.GetFullLinkByIDResponse{
			Code:   500,
			Status: fail,
			Error: &models.Err{
				Source:  "storage",
				Message: err.Error(),
			},
			Response: nil,
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
	return r.storage.Ping(ctx)
}

func (r *UserRepo) ShorteningBatchLinks(ctx context.Context, request user.ShorteningBatchLinksRequest) user.ShorteningBatchLinksResponse {
	storageMap := make(map[string]storage.Item)
	var result []user.BatchUrlsOutput
	for _, itemOfBatch := range request.Body {
		alias := utils.RandomString(sizeOfAlias)
		itemToStoreInDB := storage.Item{
			Object:     itemOfBatch.OriginalUrl,
			Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
		}
		storageMap[alias] = itemToStoreInDB
		result = append(result, user.BatchUrlsOutput{
			CorrelationID: itemOfBatch.CorrelationID,
			ShortUrl:      request.BaseURL + "/" + alias,
		})
	}
	err := r.storage.Set(ctx, storageMap)
	if err != nil {
		return user.ShorteningBatchLinksResponse{
			Code:   500,
			Status: fail,
			Error: &models.Err{
				Source:  "storage",
				Message: err.Error(),
			},
			Response: nil,
		}
	}

	if r.storage.File != nil {
		err := r.storage.File.SaveToFile(storageMap)
		if err != nil {
			return user.ShorteningBatchLinksResponse{
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

	return user.ShorteningBatchLinksResponse{
		Code:     201,
		Status:   success,
		Error:    nil,
		Response: result,
	}
}
