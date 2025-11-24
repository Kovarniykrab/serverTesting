package domain

import (
	"errors"
)

var (
	ErrUnauthorized        = errors.New("unauthorized")
	ErrNotFound            = errors.New("not found")
	ErrConflict            = errors.New("conflict")
	ErrForbidden           = errors.New("forbidden")
	ErrInternalServerError = errors.New("internal server error")
	ErrNotImplemented      = errors.New("not implemented")
	ErrBadRequest          = errors.New("bad request")
	ErrInvalidFormat       = errors.New("invalid format")
	ErrNoContent           = errors.New("no content")
	ErrTimeout             = errors.New("Timeout")
)
