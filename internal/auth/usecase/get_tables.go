package usecase

import (
	"context"
	"fmt"
)

func (u *UseCase) GetTables(ctx context.Context) error {
	err := u.postgres.GetTables(ctx)
	if err != nil {
		return fmt.Errorf("usecase get_tables error: %w", err)
	}
	return nil
}
