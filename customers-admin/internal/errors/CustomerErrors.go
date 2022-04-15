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

func NewMarshallError() *CustomerError {
	return &CustomerError{http.StatusBadRequest, "Failed to marshal the customer"}
}

func NewUserNotFoundError() *CustomerError {
	return &CustomerError{http.StatusNotFound, "User not found"}
}

func NewEmailAlreadyRegisteredError() *CustomerError {
	return &CustomerError{http.StatusBadRequest, "User with given email already exists"}
}
