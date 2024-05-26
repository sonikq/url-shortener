package user

import (
	"github.com/sonikq/url-shortener/internal/app/models"
)

// ShorteningLinkRequest -
type ShorteningLinkRequest struct {
	UserID         string
	ShorteningLink string
	RequestURL     string
	BaseURL        string
}

// ShorteningLinkResponse -
type ShorteningLinkResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *string
}
