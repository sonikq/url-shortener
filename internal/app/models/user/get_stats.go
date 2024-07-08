package user

import "github.com/sonikq/url-shortener/internal/app/models"

// GetStatsResponse -
type GetStatsResponse struct {
	Code     int
	Status   string      `json:"status"`
	Error    *models.Err `json:"error"`
	Response StatsBody
}

// StatsBody -
type StatsBody struct {
	URL   int64 `json:"urls"`
	Users int64 `json:"users"`
}
