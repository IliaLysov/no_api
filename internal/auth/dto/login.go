package dto

import (
	"fmt"
	"no_api/internal/auth/entity"
)

type LoginOutput struct {
	Token string
}

type Login struct {
	Email    string
	Password string
	IP       string
}

func (i *Login) Validate() error {
	fmt.Println(i)
	if i.Email == "" || i.Password == "" {
		return entity.ErrCredsInvalid
	}

	return nil
}
