package postgres

import (
	"context"
	"errors"
	"no_api/internal/auth/entity"

	"github.com/jackc/pgx/v5/pgconn"
)

func (p *Postgres) CreateUser(ctx context.Context, u entity.User) (id int64, err error) {
	err = p.pool.QueryRow(ctx,
		`INSERT INTO users (email, password_hash)
         VALUES ($1, $2)
         RETURNING id`,
		u.Email, u.PasswordHash,
	).Scan(&id)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, entity.ErrEmailExists
		}
		return 0, err
	}
	return id, nil
}
