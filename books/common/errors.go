package common

import (
	"net/http"
)

//CustomError custom error inteerface
type CustomError interface {
	Message() string
	Status() int
}

type customError struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e customError) Message() string {
	return e.ErrMessage
}

func (e customError) Status() int {
	return e.ErrStatus
}

//NewBadRequestError new bad request (usually bad data)
func NewBadRequestError(message string) CustomError {
	return customError{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

//NewNotFoundError not found error
func NewNotFoundError(message string) CustomError {
	return customError{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

//NewInternalServerError internal server error
func NewInternalServerError(message string, err error) CustomError {
	result := customError{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}

	return result
}
