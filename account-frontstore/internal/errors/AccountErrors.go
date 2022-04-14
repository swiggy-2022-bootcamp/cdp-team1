package errors

import "net/http"

type AccountError struct {
	Status       int
	ErrorMessage string
}

func (accountError *AccountError) Error() string {
	return accountError.ErrorMessage
}

func NewUserNotFoundError() *AccountError {
	return &AccountError{http.StatusNotFound, "User not found"}
}

func NewMalformedIdError() *AccountError {
	return &AccountError{http.StatusBadRequest, "Malformed customer id"}
}

func NewMarshallError() *AccountError {
	return &AccountError{http.StatusBadRequest, "Failed to marshal the account"}
}
