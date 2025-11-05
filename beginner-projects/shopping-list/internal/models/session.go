package models

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"
)

type Session struct {
	ID        string    `json:"id"`
	Username  string    `json:"Username"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Default session duration
const SessionDuration = 24 * time.Hour

// Common session errors
var (
	ErrSessionExpired = errors.New("session has expired")
	ErrInvalidSession = errors.New("invalid session")
)

// IsExpired checks if the session has expired
func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

// IsValid checks if session is valid (not expired)
func (s *Session) IsValid() bool {
	return !s.IsExpired()
}

// Renew extends the session expiration time
func (s *Session) Renew() {
	s.ExpiresAt = time.Now().Add(SessionDuration)
}

// NewSession creates a new session for a user
func NewSession(username string, role Role) (*Session, error) {
	// Generate a secure random session ID (32 bytes = 64 hex chars)}
	sessionID, err := generateSessionID()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Session{
		ID:        sessionID,
		Username:  username,
		Role:      role,
		CreatedAt: now,
		ExpiresAt: now.Add(SessionDuration),
	}, nil
}

// generateSessionID creates a cryptographically secure random session ID
func generateSessionID() (string, error) {
	bytes := make([]byte, 32) //32 bytes = 256 bitss of entropy
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
