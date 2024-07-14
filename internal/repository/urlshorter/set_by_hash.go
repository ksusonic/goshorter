package urlshorter

import (
	"context"
)

const setByHashQuery = `insert into shortened (
                       		hash_id,
                       		url
                       	) values (
                       	  	$1,
                       		$2
                       	) on conflict do nothing`

func (u *UrlShorter) SetURLByHash(ctx context.Context, hash, url string) error {
	_, err := u.db.Exec(
		ctx,
		setByHashQuery,
		hash,
		url,
	)
	return err
}
