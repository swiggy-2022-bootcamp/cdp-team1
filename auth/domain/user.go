package domain

import (
	"authService/errs"
)

type User struct {
	UserID   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

type UserRepositoryDB interface {
	RepoHealthCheck() *errs.AppError
	FetchByID(string) (*User, *errs.AppError)
	FetchByUsername(string) (*User, *errs.AppError)
	FetchByEmail(string) (*User, *errs.AppError)
}
