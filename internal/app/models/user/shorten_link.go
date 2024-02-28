package user

import (
	"github.com/google/uuid"
	"github.com/sonikq/url-shortener/internal/app/models"
)

type ShorteningLinkRequest struct {
	ShorteningLink string
	RequestURL     string
	BaseURL        string
}

type ShorteningLinkResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *string
}

type FileStoreURL struct {
	Uuid        uuid.UUID `json:"uuid"`
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
}
