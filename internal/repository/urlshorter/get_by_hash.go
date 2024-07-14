package urlshorter

import (
	"context"

	"github.com/ksusonic/goshorter/internal/models"
)

const getByHashQuery = `select url 
						from shortened 
						where hash_id = $1`

func (u *UrlShorter) GetURLByHash(ctx context.Context, hash string) (url string, err error) {
	row, err := u.db.Query(ctx, getByHashQuery, hash)
	if err != nil {
		return "", err
	}

	if !row.Next() {
		return "", models.ErrNotFound
	}

	return url, row.Scan(&url)
}
