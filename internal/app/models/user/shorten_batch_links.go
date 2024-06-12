package user

import "github.com/sonikq/url-shortener/internal/app/models"

// ShorteningBatchLinksRequest -
type ShorteningBatchLinksRequest struct {
	UserID     string
	RequestURL string
	BaseURL    string
	Body       []BatchUrlsInput
}

// BatchUrlsInput -
type BatchUrlsInput struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

// ShorteningBatchLinksResponse -
type ShorteningBatchLinksResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response []BatchUrlsOutput
}

// BatchUrlsOutput -
type BatchUrlsOutput struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
