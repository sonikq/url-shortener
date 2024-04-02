package user

import "github.com/sonikq/url-shortener/internal/app/models"

type ShorteningLinkJSONRequest struct {
	UserID         string
	ShorteningLink ShortenLinkJSONRequestBody
	RequestURL     string
	BaseURL        string
}

type ShortenLinkJSONRequestBody struct {
	URL string `json:"url"`
}

type ShorteningLinkJSONResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response ShortenLinkJSONResponseBody
}

type ShortenLinkJSONResponseBody struct {
	Result string `json:"result"`
}
