package handlers

import (
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
	cfg "github.com/sonikq/url-shortener/configs/app"
	"github.com/sonikq/url-shortener/internal/app/handlers/user"
	"github.com/sonikq/url-shortener/internal/app/pkg/logger"
	"github.com/sonikq/url-shortener/internal/app/pkg/middlewares"
	"github.com/sonikq/url-shortener/internal/app/services"
	"github.com/sonikq/url-shortener/internal/app/workers"
	"github.com/sonikq/url-shortener/pkg/storage"
)

// Handlers -
type Handlers struct {
	UserHandler *user.Handler
}

// Option -
type Option struct {
	Conf    cfg.Config
	Service *services.Service
	Logger  logger.Logger
	Cache   *storage.Storage
	Worker  *workers.Worker
}

// NewRouter -
func NewRouter(option Option) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(middlewares.RequestResponseLogger(option.Logger))
	router.Use(middlewares.CompressResponse(), middlewares.DecompressRequest())

	router.MaxMultipartMemory = 8 << 20

	h := &Handlers{
		UserHandler: user.New(&user.HandlerConfig{
			Service: option.Service,
			Logger:  option.Logger,
			Conf:    option.Conf,
			Worker:  option.Worker,
		}),
	}

	router.GET("/ping_url_shortener", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})

	router.POST("/", h.UserHandler.ShorteningLink)
	router.POST("/api/shorten", h.UserHandler.ShorteningLinkJSON)
	router.POST("/api/shorten/batch", h.UserHandler.ShorteningBatchLinks)

	router.GET("/:id", h.UserHandler.GetFullLinkByID)
	router.GET("/api/user/urls", h.UserHandler.GetBatchByUserID)

	router.DELETE("/api/user/urls", h.UserHandler.DeleteBatchLinks)

	router.GET("/ping", h.UserHandler.PingDB)

	router.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	router.GET("/debug/pprof/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
	router.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	router.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	router.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	router.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))

	return router
}
