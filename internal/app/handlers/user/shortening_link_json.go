package user

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/pkg/reader"
	"net/http"
	"time"
)

func (h *Handler) ShorteningLinkJSON(ctx *gin.Context) {
	bodyBytes, err := reader.GetBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error in reading body"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	var reqBody user.ShortenLinkJSONRequestBody

	unmarshalErr := json.Unmarshal(bodyBytes, &reqBody)
	if unmarshalErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request json data, cannot unmarshal into Go-struct"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}

	request := user.ShorteningLinkJSONRequest{
		ShorteningLink: reqBody,
		RequestURL:     ctx.Request.Host + ctx.Request.URL.String(),
		BaseURL:        h.config.BaseURL,
	}

	response := make(chan user.ShorteningLinkJSONResponse, 1)

	c, cancel := context.WithTimeout(ctx, time.Second*time.Duration(h.config.CtxTimeout))
	defer cancel()

	go h.service.IUserService.ShorteningLinkJSON(c, request, response)
	defer func() {
		if r := recover(); r != nil {
			h.log.Fatal("паника", logger.String("описание", "обнаружена паника"))
		}
	}()

	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			StatusKey: TimeLimitExceedErr,
		})
	case result := <-response:
		switch result.Code {
		case http.StatusCreated:
			ctx.JSON(result.Code, result.Response)
		case http.StatusConflict:
			ctx.JSON(result.Code, result.Response)
		default:
			ctx.JSON(result.Code, gin.H{
				StatusKey: result.Status,
				ErrMsgKey: result.Error.Message,
			})
		}
	}
}
