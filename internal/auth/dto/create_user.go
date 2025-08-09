package dto

import (
	"fmt"
	"no_api/internal/auth/entity"
)

type CreateUserOutput struct {
	ID int64
}

type CreateUserInput struct {
	Email    string
	Password string
}

func (i *CreateUserInput) Validate() error {
	fmt.Println(i)
	if i.Email == "" || i.Password == "" {
		return entity.ErrCredsInvalid
	}

	return nil
}
