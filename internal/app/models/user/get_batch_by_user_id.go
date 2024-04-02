package user

import "github.com/sonikq/url-shortener/internal/app/models"

type GetBatchByUserIDRequest struct {
	BaseURL string
	UserID  string
}

type GetBatchByUserIDResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response []BatchByUserID
}

type BatchByUserID struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
