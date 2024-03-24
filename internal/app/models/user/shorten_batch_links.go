package user

import "github.com/sonikq/url-shortener/internal/app/models"

type ShorteningBatchLinksRequest struct {
	RequestURL string
	BaseURL    string
	Body       []BatchUrlsInput
}

type BatchUrlsInput struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type ShorteningBatchLinksResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response []BatchUrlsOutput
}

type BatchUrlsOutput struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
