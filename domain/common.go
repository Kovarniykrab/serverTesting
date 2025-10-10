package domain

import (
	"errors"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrEmailExists     = errors.New("email already exists")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrUnauthorized    = errors.New("unauthorized access")
)
