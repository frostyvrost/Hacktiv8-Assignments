package pkg

import (
	"net/http"
	"strings"
)

type Error interface {
	Message() string
	Status() int
	Type() string
}

type ErrorResponse struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
	ErrorType    string `json:"error"`
}

func (e *ErrorResponse) Message() string {
	return e.ErrorMessage
}

func (e *ErrorResponse) Status() int {
	return e.ErrorStatus
}

func (e *ErrorResponse) Type() string {
	return e.ErrorType
}

func NewError(message string, status int, ErrorType string) Error {
	return &ErrorResponse{
		ErrorMessage: message,
		ErrorStatus:  status,
		ErrorType:    ErrorType,
	}
}

func BadRequest(message string) Error {
	return NewError(message, http.StatusBadRequest, "Bad Request")
}

func Unautorized(message string) Error {
	return NewError(message, http.StatusUnauthorized, "Unauthorized")
}

func NotFound(message string) Error {
	return NewError(message, http.StatusNotFound, "Not Found")
}

func UnprocessibleEntity(message string) Error {
	return NewError(message, http.StatusUnprocessableEntity, "Invalid Request")
}

func InternalServerError(message string) Error {
	return NewError(message, http.StatusInternalServerError, "Server Error")
}

func ParseError(err error) Error {
	if strings.Contains(err.Error(), "record not found") {
		return NotFound("Data not found")
	}

	return InternalServerError("Something went wrong")
}
