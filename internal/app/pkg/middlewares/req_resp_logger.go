package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"time"
)

func RequestResponseLogger(l logger.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		uri := ctx.Request.RequestURI
		method := ctx.Request.Method

		ctx.Next()

		l.Info("request info", logger.String("uri", uri),
			logger.String("method", method),
			logger.String("duration",
				time.Since(start).String()))

		status := ctx.Writer.Status()
		size := ctx.Writer.Size()

		l.Info("response info", logger.Int("status", status),
			logger.Int("size", size))
	}
}
