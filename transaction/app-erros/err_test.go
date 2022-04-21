package app_erros

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("Bad request")
	assert.Equal(t, err.Code, http.StatusBadRequest)
	assert.Equal(t, err.Message, "Bad request")
}

func TestNewUnexpectedError(t *testing.T) {
	err := NewUnexpectedError("Internal server error")
	assert.Equal(t, err.Code, http.StatusInternalServerError)
	assert.Equal(t, err.Message, "Internal server error")
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("Not found error")
	assert.Equal(t, err.Code, http.StatusNotFound)
	assert.Equal(t, err.Message, "Not found error")
}

func TestNewExpectationFailed(t *testing.T) {
	err := NewExpectationFailed("Expectation failed error")
	assert.Equal(t, err.Code, http.StatusExpectationFailed)
	assert.Equal(t, err.Message, "Expectation failed error")
}
