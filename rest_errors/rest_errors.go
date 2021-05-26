package rest_errors

import (
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
	errorCode  string        `json:"error"`
	causes     []interface{} `json:"causes"`
}

// Implement built-in Error go function
func (e restError) Error() string {
	return e.errorCode
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

func NewRESTError(message string, statusCode int, errorCode string, causes []interface{}) RESTError {
	return restError{
		statusCode: statusCode,
		message:    message,
		errorCode:  errorCode,
		causes:     causes,
	}
}

// NewBadRequestRESTError creates a new bad request REST error
func NewBadRequestRESTError(message string) RESTError {
	return restError{
		statusCode: http.StatusBadRequest,
		message:    message,
		errorCode:  "bad_request",
	}
}

// NewNotFoundRESTError creates a new not found REST error
func NewNotFoundRESTError(message string) RESTError {
	return restError{
		statusCode: http.StatusNotFound,
		message:    message,
		errorCode:  "not_found",
	}
}

func NewUnauthorizedRESTError(message string) RESTError {
	return restError{
		statusCode: http.StatusUnauthorized,
		message:    message,
		errorCode:  "unauthorized",
	}
}

// NewInternalServerError creates a new internal server REST error
func NewInternalServerRESTError(message string, err error) RESTError {
	return restError{
		statusCode: http.StatusInternalServerError,
		message:    message,
		errorCode:  "internal_server_error",
		causes:     []interface{}{err.Error()},
	}
}
