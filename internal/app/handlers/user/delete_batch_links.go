package user

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/pkg/auth"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/pkg/reader"
)

// DeleteBatchLinks удаляет ссылки скопом(сразу несколько штук).
//
// DELETE /api/user/urls
//
// Content-Type: application/json.
//
// В запросе - массив строк(сокращенных ссылок) [string].
func (h *Handler) DeleteBatchLinks(ctx *gin.Context) {
	userID, err := auth.VerifyUserToken(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "cant get cookie"})
		h.log.Error("userID not found, or invalid", logger.Error(err))
		return
	}

	bodyBytes, err := reader.GetBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error in reading body"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	var reqBody []string

	unmarshalErr := json.Unmarshal(bodyBytes, &reqBody)
	if unmarshalErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request json data"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	err = h.worker.DeleteURLs(reqBody, userID)
	ctx.Status(http.StatusAccepted)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while deleting links"})
		h.log.Error("worker.DeleteURLs", logger.Error(err))
		return
	}

}
