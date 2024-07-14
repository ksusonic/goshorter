package urlshorter

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlShorter struct {
	db *pgxpool.Pool
}

func NewUrlShorter(pool *pgxpool.Pool) *UrlShorter {
	return &UrlShorter{db: pool}
}
