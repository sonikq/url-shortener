package user

import "github.com/sonikq/url-shortener/internal/app/models"

// GetFullLinkByIDRequest -
type GetFullLinkByIDRequest struct {
	ShortLinkID string
}

// GetFullLinkByIDResponse -
type GetFullLinkByIDResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response *string     `json:"response"`
}
