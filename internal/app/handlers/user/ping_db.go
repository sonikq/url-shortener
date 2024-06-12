package user

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// PingDB проверка работоспособности БД.
//
// GET /ping
//
// Content-Type: text/plain.
//
// В запросе - массив строк(сокращенных ссылок) [string].
func (h *Handler) PingDB(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, CtxTimeout*time.Second)
	defer cancel()
	err := h.service.IUserService.PingDB(c)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		h.log.Error("cannot ping database")
		return
	}
	ctx.Status(http.StatusOK)
}
