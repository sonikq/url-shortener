package user

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/auth"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/pkg/reader"
	"net/http"
	"time"
)

func (h *Handler) ShorteningBatchLinks(ctx *gin.Context) {
	userID, err := auth.GetUserToken(ctx.Writer, ctx.Request)
	if err != nil {
		//ctx.Status(http.StatusBadRequest)
		h.log.Info("userID not found, or invalid", logger.Error(err))
		//return
	}

	bodyBytes, err := reader.GetBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error in reading body"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	var reqBody []user.BatchUrlsInput

	unmarshalErr := json.Unmarshal(bodyBytes, &reqBody)
	if unmarshalErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request json data, cannot unmarshal into Go-struct"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	request := user.ShorteningBatchLinksRequest{
		UserID:     userID,
		Body:       reqBody,
		RequestURL: ctx.Request.Host + ctx.Request.URL.String(),
		BaseURL:    h.config.BaseURL,
	}

	c, cancel := context.WithTimeout(ctx, CtxTimeout*time.Second)
	defer cancel()

	result := h.service.IUserService.ShorteningBatchLinks(c, request)
	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			StatusKey: TimeLimitExceedErr,
		})
	default:
		switch result.Code {
		case http.StatusCreated:
			ctx.JSON(result.Code, result.Response)
		default:
			ctx.JSON(result.Code, gin.H{
				StatusKey: result.Status,
				ErrMsgKey: result.Error.Message,
			})
		}
	}
}
