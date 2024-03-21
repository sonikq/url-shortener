package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"net/http"
	"time"
)

func (h *Handler) GetFullLinkByID(ctx *gin.Context) {
	linkID := ctx.Param("id")

	request := user.GetFullLinkByIDRequest{
		ShortLinkID: linkID,
	}

	response := make(chan user.GetFullLinkByIDResponse, 1)

	c, cancel := context.WithTimeout(ctx, time.Second*time.Duration(h.config.CtxTimeout))
	defer cancel()

	go h.service.IUserService.GetFullLinkByID(c, request, response)
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
		case http.StatusTemporaryRedirect:
			ctx.Header("Location", *result.Response)
			ctx.Status(result.Code)
		default:
			ctx.JSON(result.Code, gin.H{
				StatusKey: result.Status,
				ErrMsgKey: result.Error.Message,
			})
		}
	}
}
