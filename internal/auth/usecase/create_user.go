package usecase

import (
	"context"
	"fmt"
	"no_api/internal/auth/dto"
	"no_api/internal/auth/entity"
	"strconv"
	"time"

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

	event := entity.CreateEvent{
		ID:         strconv.Itoa(int(id)),
		OccurredAt: time.Now().UTC(),
		Type:       "user.created",
		Payload:    entity.UserCreated{Email: user.Email},
	}

	err = u.Kafka.CreateEvent(ctx, event)
	if err != nil {
		return output, err
	}

	u.email.Send(user.Email, "Account created", fmt.Sprintf("Account created with id: %d", id))

	output.ID = id
	return output, nil
}
