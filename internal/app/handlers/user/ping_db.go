package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) PingDB(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	err := h.service.IUserService.PingDB(c)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		h.log.Error("cannot ping database")
		return
	}
	ctx.Status(http.StatusOK)
}
