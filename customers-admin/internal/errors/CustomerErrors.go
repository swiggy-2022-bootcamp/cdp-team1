package errors

import "net/http"

type CustomerError struct {
	Status       int
	ErrorMessage string
}

func (customerError *CustomerError) Error() string {
	return customerError.ErrorMessage
}

func NewMalformedIdError() *CustomerError {
	return &CustomerError{http.StatusBadRequest, "Malformed customer id"}
}
