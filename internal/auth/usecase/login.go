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

func (u *UseCase) Login(ctx context.Context, input dto.Login) (dto.LoginOutput, error) {
	var output dto.LoginOutput

	user := entity.User{
		Email: input.Email,
	}

	id, hash, err := u.postgres.GetUser(ctx, user)
	if err != nil {
		return output, fmt.Errorf("u.postgres.GetUser: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(input.Password))
	if err != nil {
		return output, fmt.Errorf("password is not correct")
	}
	output.Token, err = u.JWT.CreateToken(ctx, strconv.Itoa(int(id)))
	if err != nil {
		return output, fmt.Errorf("u.jwt.CreateToken: %w", err)
	}

	event := entity.CreateEvent{
		ID:         strconv.Itoa(int(id)),
		OccurredAt: time.Now().UTC(),
		Type:       "user.logged_in",
		Payload:    entity.UserLoggedIn{Email: user.Email, IP: input.IP},
	}

	err = u.Kafka.CreateEvent(ctx, event)
	if err != nil {
		return output, err
	}

	return output, nil
}
