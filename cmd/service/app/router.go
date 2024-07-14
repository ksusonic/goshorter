package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ksusonic/goshorter/internal/service/shortener"
)

const (
	index  = "/"
	ping   = "/ping"
	toHash = "/to/:hash"

	apiShorten = "/api/shorten"
)

func setupRouter(shortenerService *shortener.Service) *gin.Engine {
	r := gin.Default()

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
