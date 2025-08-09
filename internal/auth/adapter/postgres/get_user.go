package postgres

import (
	"context"
	"fmt"
	"no_api/internal/auth/entity"
)

func (p *Postgres) GetUser(ctx context.Context, u entity.User) (id int64, hash string, err error) {
	err = p.pool.QueryRow(ctx,
		`SELECT id, password_hash FROM users WHERE email=$1`,
		u.Email,
	).Scan(&id, &hash)
	fmt.Println("GetUser id:", id)

	return id, hash, err
}
