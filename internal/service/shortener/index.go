package shortener

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ksusonic/goshorter/internal/models"
)

func (s *Service) Index(c *gin.Context) {
	c.HTML(http.StatusOK, models.IndexTmpl, nil)
}
