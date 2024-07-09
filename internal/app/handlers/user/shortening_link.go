package user

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/auth"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/pkg/reader"
)

// ShorteningLink -
func (h *Handler) ShorteningLink(ctx *gin.Context) {
	userID, err := auth.GetUserToken(ctx.Writer, ctx.Request)
	if err != nil {
		h.log.Info("userID not found, or invalid", logger.Error(err))
		setCookieErr := auth.SetUserCookie(ctx.Writer)
		if setCookieErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant set cookie"})
			h.log.Error("Set Cookie err:", logger.Error(setCookieErr))
			return
		}
	}

	body, err := reader.GetBody(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error in reading body"})
		h.log.Error("Invalid request data", logger.Error(err))
		return
	}
	request := user.ShorteningLinkRequest{
		UserID:         userID,
		ShorteningLink: string(body),
		BaseURL:        h.config.BaseURL,
	}

	c, cancel := context.WithTimeout(ctx, CtxTimeout*time.Second)
	defer cancel()

	result := h.service.IUserService.ShorteningLink(c, request)

	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			StatusKey: TimeLimitExceedErr,
		})
	default:
		switch result.Code {
		case http.StatusCreated:
			respBytes := []byte(*result.Response)
			ctx.Data(result.Code, "text/plain", respBytes)
		case http.StatusConflict:
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
