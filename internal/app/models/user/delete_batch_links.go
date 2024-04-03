package user

import "github.com/sonikq/url-shortener/internal/app/models"

type DeleteBatchLinksRequest struct {
	UserID  string
	BaseURL string
	Body    []DeleteBatchBody
}

type DeleteBatchBody struct {
	URLS []string
}

type DeleteBatchLinksResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *string     `json:"response"`
}
