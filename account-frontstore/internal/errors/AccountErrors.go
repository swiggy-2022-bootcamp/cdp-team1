package errors

import "net/http"

type AccountError struct {
	Status       int
	ErrorMessage string
}

func (accountError *AccountError) Error() string {
	return accountError.ErrorMessage
}

func NewMalformedIdError() *AccountError {
	return &AccountError{http.StatusBadRequest, "Malformed customer id"}
}
