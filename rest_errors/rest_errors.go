package rest_errors

import (
	"fmt"
	"net/http"
)

type RESTError interface {
	StatusCode() int
	Message() string
	Error() string
	Causes() []interface{}
}

// restError struct is an custom REST error is used for entire microservice
type restError struct {
	statusCode int           `json:"status_code"`
	message    string        `json:"message"`
	error      string        `json:"error"`
	causes     []interface{} `json:"causes"`
}

// Implement built-in Error go function
func (e restError) Error() string {
	return fmt.Sprintf("message: %s - status_code: %d - error: %s - causes: %v", e.message, e.statusCode, e.error, e.causes)
}

func (e restError) StatusCode() int {
	return e.statusCode
}

func (e restError) Message() string {
	return e.message
}

func (e restError) Causes() []interface{} {
	return e.causes
}

func NewRESTError(message string, statusCode int, err string, causes []interface{}) RESTError {
	return restError{
		statusCode: statusCode,
		message:    message,
		error:      err,
		causes:     causes,
	}
}

// NewBadRequestRESTError creates a new bad request REST error
func NewBadRequestRESTError(message string) RESTError {
	return restError{
		statusCode: http.StatusBadRequest,
		message:    message,
		error:      "bad_request",
	}
}

// NewNotFoundRESTError creates a new not found REST error
func NewNotFoundRESTError(message string) RESTError {
	return restError{
		statusCode: http.StatusNotFound,
		message:    message,
		error:      "not_found",
	}
}

func NewUnauthorizedRESTError(message string) RESTError {
	return restError{
		statusCode: http.StatusUnauthorized,
		message:    message,
		error:      "unauthorized",
	}
}

// NewInternalServerError creates a new internal server REST error
func NewInternalServerRESTError(message string, err error) RESTError {
	return restError{
		statusCode: http.StatusInternalServerError,
		message:    message,
		error:      "internal_server_error",
		causes:     []interface{}{err.Error()},
	}
}
