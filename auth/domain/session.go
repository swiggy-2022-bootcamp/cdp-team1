package domain

import (
	"authService/errs"
)

type Session struct {
	TokenSignature string `json:"token_signature,omitempty"`
	UserID         string `json:"user_id,omitempty"`
	Role           string `json:"role,omitempty"`
	IsActive       bool   `json:"is_active,omitempty"`
}

func NewSession(tokenSign string, userID string, role string, isActive bool) *Session {
	return &Session{
		TokenSignature: tokenSign,
		UserID:         userID,
		Role:           role,
		IsActive:       isActive,
	}
}

type AuthRepositoryDB interface {
	RepoHealthCheck() *errs.AppError
	SaveSession(Session) *errs.AppError
	FetchSessionByTokenSign(string) (*Session, *errs.AppError)
	UpdateSessionStatusByTokenSign(string) *errs.AppError
}
