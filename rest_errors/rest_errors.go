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
	ErrorStatusCode int           `json:"status_code"`
	ErrorMessage    string        `json:"message"`
	ErrorCode       string        `json:"error"`
	ErrorCauses     []interface{} `json:"causes"`
}

// Implement built-in Error go function
func (e *restError) Error() string {
	return e.ErrorCode
}

func (e *restError) StatusCode() int {
	return e.ErrorStatusCode
}

func (e *restError) Message() string {
	return e.ErrorMessage
}

func (e *restError) Causes() []interface{} {
	return e.ErrorCauses
}

func NewRESTError(message string, statusCode int, errCode string, causes []interface{}) RESTError {
	return &restError{
		ErrorStatusCode: statusCode,
		ErrorMessage:    message,
		ErrorCode:       errCode,
		ErrorCauses:     causes,
	}
}

// NewBadRequestRESTError creates a new bad request REST error
func NewBadRequestRESTError(message string) RESTError {
	return &restError{
		ErrorStatusCode: http.StatusBadRequest,
		ErrorMessage:    message,
		ErrorCode:       "bad_request",
	}
}

// NewNotFoundRESTError creates a new not found REST error
func NewNotFoundRESTError(message string) RESTError {
	return &restError{
		ErrorStatusCode: http.StatusNotFound,
		ErrorMessage:    message,
		ErrorCode:       "not_found",
	}
}

func NewUnauthorizedRESTError(message string) RESTError {
	return &restError{
		ErrorStatusCode: http.StatusUnauthorized,
		ErrorMessage:    message,
		ErrorCode:       "unauthorized",
	}
}

// NewInternalServerError creates a new internal server REST error
func NewInternalServerRESTError(message string, err error) RESTError {
	return &restError{
		ErrorStatusCode: http.StatusInternalServerError,
		ErrorMessage:    message,
		ErrorCode:       "internal_server_error",
		ErrorCauses:     []interface{}{err.Error()},
	}
}
