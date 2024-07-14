package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	us "github.com/ksusonic/goshorter/internal/repository/urlshorter"
)

type Repository struct {
	*us.UrlShorter
}

func NewRepository(ctx context.Context, dsn string) (*Repository, func()) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("pgx connect: %+v", err)
	}

	return &Repository{
		UrlShorter: us.NewUrlShorter(pool),
	}, pool.Close
}
