package urlshorter

import (
	"context"

	"github.com/ksusonic/goshorter/internal/models"
)

const getByHashQuery = `select url 
						from shortened 
						where hash_id = $1`

func (u *URLShorter) GetURLByHash(ctx context.Context, hash string) (string, error) {
	row, err := u.db.Query(ctx, getByHashQuery, hash)
	if err != nil {
		return "", err
	}

	if !row.Next() {
		return "", models.ErrNotFound
	}

	var url string
	return url, row.Scan(&url)
}
