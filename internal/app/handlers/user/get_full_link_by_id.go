package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/models/user"
	"net/http"
	"time"
)

func (h *Handler) GetFullLinkByID(ctx *gin.Context) {
	linkID := ctx.Param("id")

	request := user.GetFullLinkByIDRequest{
		ShortLinkID: linkID,
	}

	c, cancel := context.WithTimeout(ctx, CtxTimeout*time.Second)
	defer cancel()

	result := h.service.IUserService.GetFullLinkByID(c, request)
	select {
	case <-c.Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			StatusKey: TimeLimitExceedErr,
		})
	default:
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
