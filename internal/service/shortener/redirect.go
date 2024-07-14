package shortener

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ksusonic/goshorter/internal/models"
)

type Redirect struct {
	Hash string `uri:"hash" binding:"required"`
}

func (s *Service) Redirect(c *gin.Context) {
	var redirect Redirect
	if err := c.ShouldBindUri(&redirect); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	redirectURL, err := s.repo.GetURLByHash(c.Request.Context(), redirect.Hash)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.HTML(http.StatusNotFound, models.NotFoundTmpl, nil)
			return
		}

		s.log.Printf("repo.GetURLByHash: %+v", err)
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
