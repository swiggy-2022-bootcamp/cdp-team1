package errs

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//TestAppError_Error ..
func TestAppError_Error(t *testing.T) {
	err := NewBadRequestError("")
	errObj := err.Error()
	assert.Error(t, errObj)
}

//TestAppError_AsMessage ..
func TestAppError_AsMessage(t *testing.T) {
	err := NewBadRequestError("error")
	errMsg := err.AsMessage()
	assert.Equal(t, errMsg.Message, "error")
}

//TestNewNotFoundError ..
func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("Not found")
	assert.Equal(t, err.Code, http.StatusNotFound)
	assert.Equal(t, err.Message, "Not found")
}

//TestNewBadRequestError ..
func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("Bad request")
	assert.Equal(t, err.Code, http.StatusBadRequest)
	assert.Equal(t, err.Message, "Bad request")
}

//TestNewUnexpectedError ..
func TestNewUnexpectedError(t *testing.T) {
	err := NewUnexpectedError("Internal server error")
	assert.Equal(t, err.Code, http.StatusInternalServerError)
	assert.Equal(t, err.Message, "Internal server error")
}
