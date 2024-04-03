package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/auth"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"net/http"
	"time"
)

func (h *Handler) GetBatchByUserID(ctx *gin.Context) {
	userID, err := auth.VerifyUserToken(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "cant get cookie"})
		h.log.Error("userID not found, or invalid", logger.Error(err))
		return
	}

	request := user.GetBatchByUserIDRequest{
		UserID:  userID,
		BaseURL: h.config.BaseURL,
	}

	c, cancel := context.WithTimeout(ctx, CtxTimeout*time.Minute)
	defer cancel()

	result := h.service.IUserService.GetBatchByUserID(c, request)
	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			StatusKey: TimeLimitExceedErr,
		})
	default:
		switch result.Code {
		case http.StatusNoContent:
			ctx.JSON(result.Code, gin.H{
				StatusKey: result.Code,
				ErrMsgKey: "no content found",
			})
		case http.StatusOK:
			ctx.JSON(result.Code, result.Response)
		default:
			ctx.JSON(result.Code, gin.H{
				StatusKey: result.Status,
				ErrMsgKey: result.Error.Message,
			})
		}
	}
}
