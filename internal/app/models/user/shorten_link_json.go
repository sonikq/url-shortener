package user

import "github.com/sonikq/url-shortener/internal/app/models"

// ShorteningLinkJSONRequest -
type ShorteningLinkJSONRequest struct {
	UserID         string
	ShorteningLink ShortenLinkJSONRequestBody
	RequestURL     string
	BaseURL        string
}

// ShortenLinkJSONRequestBody -
type ShortenLinkJSONRequestBody struct {
	URL string `json:"url"`
}

// ShorteningLinkJSONResponse -
type ShorteningLinkJSONResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response ShortenLinkJSONResponseBody
}

// ShortenLinkJSONResponseBody -
type ShortenLinkJSONResponseBody struct {
	Result string `json:"result"`
}
