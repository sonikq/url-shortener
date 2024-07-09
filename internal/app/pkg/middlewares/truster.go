package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sonikq/url-shortener/configs/app"
	"net"
	"net/http"
)

// Truster middleware для проверки находится ли IP-адрес клиента в доверенной подсети
func Truster(conf app.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Получение IP-адреса клиента из заголовка X-Real-IP
		clientIP := ctx.ClientIP()

		// Парсинг подсети
		_, trustedIPNet, err := net.ParseCIDR(conf.TrustedSubnet)
		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"описание ошибки": "Не указан IP-адрес клиента"})
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		// Парсинг IP-адреса клиента
		clientAddr := net.ParseIP(clientIP)
		if clientAddr == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"описание ошибки": "Невозможно определить адрес клиента"})
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Проверка, находится ли IP-адрес клиента в доверенной подсети
		if !trustedIPNet.Contains(clientAddr) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"описание ошибки": "Доступ запрещен"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Next()
	}
}
