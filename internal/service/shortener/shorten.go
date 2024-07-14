package shortener

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/ksusonic/goshorter/internal/utils"
)

type shortenRequest struct {
	URL string `json:"url"`
}

func (s *Service) Shorten(c *gin.Context) {
	var request shortenRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err := url.ParseRequestURI(request.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	hash := utils.GenerateHash(request.URL)

	err = s.repo.SetURLByHash(c.Request.Context(), hash, request.URL)
	if err != nil {
		s.log.Printf("repo.SetURLByHash: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	shortUrl := s.shortURLPrefix + hash

	c.JSON(http.StatusOK, gin.H{"shortened_url": shortUrl})
}
