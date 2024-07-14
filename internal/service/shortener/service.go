package shortener

import (
	"context"
	"log"
)

type Service struct {
	shortURLPrefix string

	repo Repository
	log  *log.Logger
}

func NewService(
	shortURLPrefix string,
	repo Repository,
	log *log.Logger,
) *Service {
	return &Service{
		shortURLPrefix: shortURLPrefix,
		repo:           repo,
		log:            log,
	}
}

type Repository interface {
	GetURLByHash(ctx context.Context, hash string) (string, error)
	SetURLByHash(ctx context.Context, hash, url string) error
}
