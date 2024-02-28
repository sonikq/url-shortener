package user

import (
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
