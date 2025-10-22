package domain

import (
	"errors"

	"github.com/valyala/fasthttp"
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
)

type MessageError struct {
	Message    string
	StatusCode int
}

func (e *MessageError) Error() string {
	return e.Message
}

func BadRequest(err error) error {
	if err == nil {
		err = ErrBadRequest
	}
	return &MessageError{Message: err.Error(), StatusCode: fasthttp.StatusBadRequest}
}

func Unauthorized(err error) error {
	if err == nil {
		err = ErrUnauthorized
	}
	return &MessageError{Message: err.Error(), StatusCode: fasthttp.StatusUnauthorized}
}

func Forbidden(err error) error {
	if err == nil {
		err = ErrForbidden
	}
	return &MessageError{Message: err.Error(), StatusCode: fasthttp.StatusForbidden}
}

func NotFound(err error) error {
	if err == nil {
		err = ErrNotFound
	}
	return &MessageError{Message: err.Error(), StatusCode: fasthttp.StatusNotFound}
}

func Conflict(err error) error {
	if err == nil {
		err = ErrConflict
	}
	return &MessageError{Message: err.Error(), StatusCode: fasthttp.StatusConflict}
}

func Internal(err error) error {
	if err == nil {
		err = ErrInternalServerError
	}
	return &MessageError{Message: err.Error(), StatusCode: fasthttp.StatusInternalServerError}
}

func NoContent(err error) error {
	if err == nil {
		err = ErrNoContent
	}

	return &MessageError{Message: err.Error(), StatusCode: fasthttp.StatusNoContent}
}
