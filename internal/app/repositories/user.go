package repositories

import (
	"context"
	"github.com/sonikq/url-shortener/internal/app/models"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/utils"
	"github.com/sonikq/url-shortener/pkg/storage"
	"net/http"
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

	mapToStore := utils.ConvertDataToStore(alias, request.ShorteningLink, request.UserID)

	err := r.storage.Set(ctx, mapToStore)
	if err != nil {
		if err == storage.ErrAlreadyExists {
			conflictShortURL, noShortURLErr := r.storage.GetShortURL(ctx, request.ShorteningLink)
			if noShortURLErr != nil {
				return user.ShorteningLinkResponse{
					Code:   http.StatusInternalServerError,
					Status: fail,
					Error: &models.Err{
						Source:  "storage",
						Message: noShortURLErr.Error(),
					},
					Response: nil,
				}
			}
			conflictURL := request.BaseURL + "/" + conflictShortURL
			return user.ShorteningLinkResponse{
				Code:     http.StatusConflict,
				Status:   success,
				Error:    nil,
				Response: &conflictURL,
			}
		}
		return user.ShorteningLinkResponse{
			Code:   http.StatusInternalServerError,
			Status: fail,
			Error: &models.Err{
				Source:  "storage",
				Message: err.Error(),
			},
			Response: nil,
		}
	}

	if r.storage.File != nil {
		err = r.storage.File.SaveToFile(mapToStore)
		if err != nil {
			return user.ShorteningLinkResponse{
				Code:   http.StatusInternalServerError,
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
		Code:     http.StatusCreated,
		Status:   success,
		Error:    nil,
		Response: &result,
	}
}

func (r *UserRepo) ShorteningLinkJSON(ctx context.Context, request user.ShorteningLinkJSONRequest) user.ShorteningLinkJSONResponse {
	alias := utils.RandomString(sizeOfAlias)

	result := request.BaseURL + "/" + alias

	mapToStore := utils.ConvertDataToStore(alias, request.ShorteningLink.URL, request.UserID)

	err := r.storage.Set(ctx, mapToStore)
	if err != nil {
		if err == storage.ErrAlreadyExists {
			conflictShortURL, noShortURLErr := r.storage.GetShortURL(ctx, request.ShorteningLink.URL)
			if noShortURLErr != nil {
				return user.ShorteningLinkJSONResponse{
					Code:   http.StatusInternalServerError,
					Status: fail,
					Error: &models.Err{
						Source:  "storage, get_short_url",
						Message: noShortURLErr.Error(),
					},
					Response: user.ShortenLinkJSONResponseBody{},
				}
			}
			conflictURL := request.BaseURL + "/" + conflictShortURL
			return user.ShorteningLinkJSONResponse{
				Code:     http.StatusConflict,
				Status:   success,
				Error:    nil,
				Response: user.ShortenLinkJSONResponseBody{Result: conflictURL},
			}
		}
		return user.ShorteningLinkJSONResponse{
			Code:   http.StatusInternalServerError,
			Status: fail,
			Error: &models.Err{
				Source:  "storage, set_value",
				Message: err.Error(),
			},
			Response: user.ShortenLinkJSONResponseBody{},
		}
	}

	if r.storage.File != nil {
		err = r.storage.File.SaveToFile(mapToStore)
		if err != nil {
			return user.ShorteningLinkJSONResponse{
				Code:   http.StatusInternalServerError,
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
		Code:     http.StatusCreated,
		Status:   success,
		Error:    nil,
		Response: user.ShortenLinkJSONResponseBody{Result: result},
	}
}

func (r *UserRepo) GetFullLinkByID(ctx context.Context, request user.GetFullLinkByIDRequest) user.GetFullLinkByIDResponse {
	fullLink, err := r.storage.Get(ctx, request.ShortLinkID)
	if err != nil {
		return user.GetFullLinkByIDResponse{
			Code:   http.StatusInternalServerError,
			Status: fail,
			Error: &models.Err{
				Source:  "storage",
				Message: err.Error(),
			},
			Response: nil,
		}
	}

	return user.GetFullLinkByIDResponse{
		Code:     http.StatusTemporaryRedirect,
		Status:   success,
		Error:    nil,
		Response: &fullLink,
	}
}

func (r *UserRepo) GetBatchByUserID(ctx context.Context, request user.GetBatchByUserIDRequest) user.GetBatchByUserIDResponse {
	var result []user.BatchByUserID

	batch, err := r.storage.GetBatchByUserID(ctx, request.UserID)
	if err != nil {
		return user.GetBatchByUserIDResponse{
			Code:   http.StatusInternalServerError,
			Status: fail,
			Error: &models.Err{
				Source:  "storage",
				Message: err.Error(),
			},
			Response: nil,
		}
	}

	if len(batch) == 0 {
		return user.GetBatchByUserIDResponse{
			Code:     http.StatusNoContent,
			Status:   fail,
			Response: nil,
		}
	}

	for key, value := range batch {
		result = append(result, user.BatchByUserID{
			ShortURL:    request.BaseURL + "/" + key,
			OriginalURL: value.Object,
		})
	}

	return user.GetBatchByUserIDResponse{
		Code:     http.StatusTemporaryRedirect,
		Status:   success,
		Error:    nil,
		Response: result,
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
			Object:     itemOfBatch.OriginalURL,
			Expiration: time.Now().Add(10 * time.Minute).UnixNano(),
			UserID:     request.UserID,
		}
		storageMap[alias] = itemToStoreInDB
		result = append(result, user.BatchUrlsOutput{
			CorrelationID: itemOfBatch.CorrelationID,
			ShortURL:      request.BaseURL + "/" + alias,
		})
	}
	err := r.storage.Set(ctx, storageMap)
	if err != nil {
		return user.ShorteningBatchLinksResponse{
			Code:   http.StatusInternalServerError,
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
				Code:   http.StatusInternalServerError,
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
		Code:     http.StatusCreated,
		Status:   success,
		Error:    nil,
		Response: result,
	}
}
