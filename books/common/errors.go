package common

import (
	"net/http"
)

//Error custom error inteerface
type Error interface {
	Message() string
	Status() int
}

//Error custom error struct
type error struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e error) Message() string {
	return e.ErrMessage
}

func (e error) Status() int {
	return e.ErrStatus
}

//NewBadRequestError new bad request (usually bad data)
func NewBadRequestError(message string) Error {
	return error{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

//NewNotFoundError item not found
func NewNotFoundError(message string) Error {
	return error{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}
