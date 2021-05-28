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
	assert.EqualValues(t, "forbidden", err.Error())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "Error causes", err.Causes()[0])
}

func TestNewRESTErrorFromBytesInvalidJSONBytes(t *testing.T) {
	jsonBytes := []byte(`{"status_code":"404","message":"No access token found","error":"not_found"}`)
	restErr, err := NewRESTErrorFromBytes(jsonBytes)
	assert.Nil(t, restErr)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid JSON bytes", err.Error())
}

func TestNewRESTErrorFromBytesNoError(t *testing.T) {
	jsonBytes := []byte(`{"status_code":404,"message":"No access token found","error":"not_found","causes":["Error causes"]}`)
	restErr, err := NewRESTErrorFromBytes(jsonBytes)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, restErr)
	assert.EqualValues(t, http.StatusNotFound, restErr.StatusCode())
	assert.EqualValues(t, "No access token found", restErr.Message())
	assert.EqualValues(t, "not_found", restErr.Error())
	assert.NotNil(t, restErr.Causes())
	assert.EqualValues(t, 1, len(restErr.Causes()))
	assert.EqualValues(t, "Error causes", restErr.Causes()[0])
}

func TestNewInternalServerRESTError(t *testing.T) {
	err := NewInternalServerRESTError("This is the test message", errors.New("Error causes"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "internal_server_error", err.Error())

	assert.NotNil(t, err.Causes())
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "Error causes", err.Causes()[0])
}

func TestNewNotFoundRESTError(t *testing.T) {
	err := NewNotFoundRESTError("This is the test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "not_found", err.Error())
	assert.Nil(t, err.Causes())
}

func TestNewBadRequestRESTError(t *testing.T) {
	err := NewBadRequestRESTError("This is the test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "bad_request", err.Error())
	assert.Nil(t, err.Causes())
}

func TestNewUnauthorizedRESTError(t *testing.T) {
	err := NewUnauthorizedRESTError("This is the test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode())
	assert.EqualValues(t, "This is the test message", err.Message())
	assert.EqualValues(t, "unauthorized", err.Error())
	assert.Nil(t, err.Causes())
}
