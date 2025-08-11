package usecase

import (
	"context"
	"no_api/internal/auth/entity"
)

type Postgres interface {
	GetTables(ctx context.Context) (err error)
	CreateUser(ctx context.Context, u entity.User) (id int64, err error)
	GetUser(ctx context.Context, u entity.User) (id int64, hash string, err error)
}

type JWT interface {
	CreateToken(context.Context, string) (string, error)
	Verify(string) (string, error)
}

type Email interface {
	Send(to, subject, body string) error
}

type UseCase struct {
	postgres Postgres
	JWT      JWT
	email    Email
}

func New(p Postgres, jwt JWT, email Email) *UseCase {
	return &UseCase{postgres: p, JWT: jwt, email: email}
}
