package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRESTError(t *testing.T) {
	err := NewRESTError("This is the test message", http.StatusForbidden, "forbidden", []interface{}{"Error causes"})
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusForbidden, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "message: This is the test message - status_code: 403 - error: forbidden - causes: [Error causes]", err.Error())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "Error causes", err.Causes()[0])
}

func TestNewInternalServerRESTError(t *testing.T) {
	err := NewInternalServerRESTError("This is the test message", errors.New("Error causes"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "message: This is the test message - status_code: 500 - error: internal_server_error - causes: [Error causes]", err.Error())

	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "Error causes", err.Causes()[0])
}

func TestNewNotFoundRESTError(t *testing.T) {
	err := NewNotFoundRESTError("This is the test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "message: This is the test message - status_code: 404 - error: not_found - causes: []", err.Error())
}

func TestNewBadRequestRESTError(t *testing.T) {
	err := NewBadRequestRESTError("This is the test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "message: This is the test message - status_code: 400 - error: bad_request - causes: []", err.Error())
}

func TestNewUnauthorizedRESTError(t *testing.T) {
	err := NewUnauthorizedRESTError("This is the test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "message: This is the test message - status_code: 401 - error: unauthorized - causes: []", err.Error())
}
