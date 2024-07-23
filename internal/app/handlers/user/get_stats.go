package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetStats получение статистики о сокращенных URL и пользователях.
//
// GET /api/internal/stats
//
// Content-Type: text/plain.
func (h *Handler) GetStats(ctx *gin.Context) {

	c, cancel := context.WithTimeout(ctx, CtxTimeout*time.Second)
	defer cancel()

	result := h.service.IUserService.GetStats(c)
	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			StatusKey: TimeLimitExceedErr,
		})
	default:
		switch result.Code {
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
