package usecase

import (
	"context"
	"fmt"
	"no_api/internal/auth/dto"
	"no_api/internal/auth/entity"

	"golang.org/x/crypto/bcrypt"
)

func (u *UseCase) CreateUser(ctx context.Context, input dto.CreateUserInput) (dto.CreateUserOutput, error) {
	var output dto.CreateUserOutput

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := entity.User{
		Email:        input.Email,
		PasswordHash: string(hash),
	}
	id, err := u.postgres.CreateUser(ctx, user)
	if err != nil {
		return output, fmt.Errorf("u.postgres.CreateUser: %w", err)
	}

	u.email.Send(user.Email, "Account created", fmt.Sprintf("Account created with id: %d", id))

	output.ID = id
	return output, nil
}
