package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/reader"
	"log"
	"net/http"
	"time"
)

func (h *Handler) ShorteningLink(ctx *gin.Context) {
	body, err := reader.GetBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error in reading body"})
		return
	}
	request := user.ShorteningLinkRequest{
		ShorteningLink: string(body),
		RequestURL:     ctx.Request.Host + ctx.Request.URL.String(),
		BaseURL:        h.config.BaseURL,
	}

	response := make(chan user.ShorteningLinkResponse, 1)

	c, cancel := context.WithTimeout(ctx, time.Second*time.Duration(h.config.CtxTimeout))
	defer cancel()

	go h.service.IUserService.ShorteningLink(request, response)
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("обнаружена паника")
		}
	}()

	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			StatusKey: TimeLimitExceedErr,
		})
	case result := <-response:
		switch result.Code {
		case 201:
			respBytes := []byte(*result.Response)
			ctx.Data(result.Code, "text/plain", respBytes)
		default:
			ctx.JSON(result.Code, gin.H{
				StatusKey: result.Status,
				ErrMsgKey: result.Error.Message,
			})
		}
	}
}
