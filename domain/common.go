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

type MessageError struct {
	Message    string
	StatusCode int
}

func (e MessageError) Error() string {
	return e.Message
}

func BadRequest(err error) error {
	return MessageError{Message: err.Error(), StatusCode: 400}
}

func Unauthorized(err error) error {
	return MessageError{Message: err.Error(), StatusCode: 401}
}

func Forbidden(err error) error {
	return MessageError{Message: err.Error(), StatusCode: 403}
}

func NotFound(err error) error {
	return MessageError{Message: err.Error(), StatusCode: 404}
}

func Conflict(err error) error {
	return MessageError{Message: err.Error(), StatusCode: 409}
}

func Internal(err error) error {
	return MessageError{Message: err.Error(), StatusCode: 500}
}
