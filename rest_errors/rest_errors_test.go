package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerRESTError(t *testing.T) {
	err := NewInternalServerRESTError("This is the test message", errors.New("Error causes"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "This is the test message", err.Message)
	assert.EqualValues(t, "INTERNAL SERVER ERROR", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "Error causes", err.Causes[0])
}

func TestNewNotFoundRESTError(t *testing.T) {
	err := NewNotFoundRESTError("This is the test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "This is the test message", err.Message)
	assert.EqualValues(t, "NOT FOUND", err.Error)
}

func TestNewBadRequestRESTError(t *testing.T) {
	err := NewBadRequestRESTError("This is the test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode)
	assert.EqualValues(t, "This is the test message", err.Message)
	assert.EqualValues(t, "BAD REQUEST", err.Error)
}
