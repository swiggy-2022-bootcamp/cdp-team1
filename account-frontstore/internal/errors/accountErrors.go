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

func NewEmailAlreadyRegisteredError() *AccountError {
	return &AccountError{http.StatusBadRequest, "User with given email already exists"}
}

func NewNotFoundError(ErrorMessage string) *AccountError {
	return &AccountError{
		ErrorMessage: ErrorMessage,
		Status:       http.StatusNotFound,
	}
}

func NewUnexpectedError(ErrorMessage string) *AccountError {
	return &AccountError{
		ErrorMessage: ErrorMessage,
		Status:       http.StatusInternalServerError,
	}
}

func NewValidationError(ErrorMessage string) *AccountError {
	return &AccountError{
		ErrorMessage: ErrorMessage,
		Status:       http.StatusBadRequest,
	}
}

func NewAuthenticationError(ErrorMessage string) *AccountError {
	return &AccountError{
		ErrorMessage: ErrorMessage,
		Status:       http.StatusUnauthorized,
	}
}

func NewAuthorizationError(ErrorMessage string) *AccountError {
	return &AccountError{
		ErrorMessage: ErrorMessage,
		Status:       http.StatusForbidden,
	}
}
