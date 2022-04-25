package db

import (
	"authService/domain"
	"time"
)

type Session struct {
	TokenSignature string    `json:"token_signature"`
	UserID         string    `json:"user_id"`
	Role           string    `json:"role"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewDBSession(domSession domain.Session) *Session {
	return &Session{
		TokenSignature: domSession.TokenSignature,
		UserID:         domSession.UserID,
		Role:           domSession.Role,
		IsActive:       domSession.IsActive,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (s Session) DomainSession() *domain.Session {
	return &domain.Session{
		TokenSignature: s.TokenSignature,
		UserID:         s.UserID,
		Role:           s.Role,
		IsActive:       s.IsActive,
	}
}
