package entity

import "errors"

var (
	ErrCredsInvalid = errors.New("invalid credentials")
	ErrEmailExists  = errors.New("email already exists")
)
