package app_errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAppError_Error(t *testing.T) {
	err := NewBadRequestError("")
	errObj := err.Error()
	assert.Error(t, errObj)
}

func TestAppError_AsMessage(t *testing.T) {
	err := NewBadRequestError("error")
	errMsg := err.AsMessage()
	assert.Equal(t, errMsg.Message, "error")
}
func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("Not found")
	assert.Equal(t, err.Code, http.StatusNotFound)
	assert.Equal(t, err.Message, "Not found")
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("Bad request")
	assert.Equal(t, err.Code, http.StatusBadRequest)
	assert.Equal(t, err.Message, "Bad request")
}

func TestNewConflictRequestError(t *testing.T) {
	err := NewConflictRequestError("Conflict request")
	assert.Equal(t, err.Code, http.StatusConflict)
	assert.Equal(t, err.Message, "Conflict request")
}

func TestNewUnexpectedError(t *testing.T) {
	err := NewUnexpectedError("Internal server error")
	assert.Equal(t, err.Code, http.StatusInternalServerError)
	assert.Equal(t, err.Message, "Internal server error")
}

func TestNewRequestNotAcceptedError(t *testing.T) {
	err := NewRequestNotAcceptedError("Request not accepted")
	assert.Equal(t, err.Code, http.StatusNotAcceptable)
	assert.Equal(t, err.Message, "Request not accepted")
}

func TestNewAuthenticationError(t *testing.T) {
	err := NewAuthenticationError("Unauthorized access")
	assert.Equal(t, err.Code, http.StatusUnauthorized)
	assert.Equal(t, err.Message, "Unauthorized access")
}
