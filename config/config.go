package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func New() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("godotenv.Load: %w", err)
	}

	return nil
}
