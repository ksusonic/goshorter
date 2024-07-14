package urlshorter

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLShorter struct {
	db *pgxpool.Pool
}

func NewURLShorter(pool *pgxpool.Pool) *URLShorter {
	return &URLShorter{db: pool}
}
