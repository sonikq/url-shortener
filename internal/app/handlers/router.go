package handlers

import (
	"github.com/gin-gonic/gin"
	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/handlers/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/pkg/middlewares"
	"github.com/sonikq/url-shortener/internal/app/services"
	"github.com/sonikq/url-shortener/pkg/cache"
	"net/http"
)

type Handlers struct {
	UserHandler *user.Handler
}

type Option struct {
	Conf    cfg.Config
	Service *services.Service
	Logger  logger.Logger
	Cache   *cache.Cache
}

func NewRouter(option Option) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	//router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.RequestResponseLogger(option.Logger))

	router.MaxMultipartMemory = 8 << 20

	h := &Handlers{
		UserHandler: user.New(&user.HandlerConfig{
			Service: option.Service,
			Logger:  option.Logger,
			Conf:    option.Conf,
		}),
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})

	router.POST("/", h.UserHandler.ShorteningLink)

	router.GET("/:id", h.UserHandler.GetFullLinkByID)

	return router
}
