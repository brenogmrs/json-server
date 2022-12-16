package utils

import "net/http"

type CustomError struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *CustomError) Error() string {
	return r.Message
}

func NewCustomError(message, err string, code int, causes []Causes) *CustomError {
	return &CustomError{
		Message: message,
		Err:     err,
		Code:    code,
	}
}

func NewBadRequestError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "BAD_REQUEST",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "INTERNAL_SERVER_ERROR",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "NOT_FOUND",
		Code:    http.StatusNotFound,
	}
}
