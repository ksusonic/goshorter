package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksusonic/ginzap"
	"go.uber.org/zap"

	"github.com/ksusonic/goshorter/internal/service/shortener"
)

const (
	index  = "/"
	toHash = "/:hash"
	ping   = "/ping"

	apiShorten = "/api/shorten"
)

func setupRouter(log *zap.Logger, shortenerService *shortener.Service) *gin.Engine {
	r := gin.New()

	r.Use(ginzap.GinzapWithConfig(log, newGinZapConfig()))
	r.Use(ginzap.RecoveryWithZap(log, true))

	r.LoadHTMLGlob("frontend/templates/*")
	r.Static("/static", "frontend/static")

	r.GET(ping, func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET(index, shortenerService.Index)
	r.POST(apiShorten, shortenerService.Shorten)
	r.GET(toHash, shortenerService.Redirect)

	return r
}
