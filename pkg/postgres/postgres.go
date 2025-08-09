package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// type Config struct {
// 	User     string `envconfig:"POSTGRES_USER"     required:"true"`
// 	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
// 	Port     string `envconfig:"POSTGRES_PORT"     required:"true"`
// 	Host     string `envconfig:"POSTGRES_HOST"     required:"true"`
// 	DBName   string `envconfig:"POSTGRES_DB_NAME"  required:"true"`
// }

type Pool struct {
	*pgxpool.Pool
}

func New(ctx context.Context) (*Pool, error) {
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &Pool{Pool: pool}, nil
}
