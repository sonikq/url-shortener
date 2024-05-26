package user

import "github.com/sonikq/url-shortener/internal/app/models"

// DeleteBatchLinksRequest -
type DeleteBatchLinksRequest struct {
	UserID  string
	BaseURL string
	Body    []DeleteBatchBody
}

// DeleteBatchBody -
type DeleteBatchBody struct {
	URLS []string
}

// DeleteBatchLinksResponse -
type DeleteBatchLinksResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *string     `json:"response"`
}
