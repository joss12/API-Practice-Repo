package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

// User represents a system user
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	Role         Role   `json:"role"`
	DisplayName  string `json:"display_name"`
}

// Common user errors
var (
	ErrInvalidUsername = errors.New("username cannot be empty")
	ErrInvalidPassword = errors.New("password cannot be empty")
	ErrInvalidRole     = errors.New("invalid role")
)

// IsAdmin checks if user has admin privileges
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// IsUser checks if user has regular user privileges
func (u *User) IsUser() bool {
	return u.Role == RoleUser
}

// CheckPassword verifies if the provided password matches the Password Hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

// NewUser creates a new user with hashed password
func NewUser(username, password, displayName string, role Role) (*User, error) {
	if username == "" {
		return nil, ErrInvalidUsername
	}

	if password == "" {
		return nil, ErrInvalidPassword
	}

	if role != RoleUser && role != RoleAdmin {
		return nil, ErrInvalidRole
	}

	// Hash the password (cost 10 is good balance of security/performance)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	if displayName == "" {
		displayName = username
	}

	return &User{
		Username:     username,
		PasswordHash: string(hash),
		Role:         role,
		DisplayName:  displayName,
	}, nil
}

func ValidateRole(role string) bool {
	return role == string(RoleUser) || role == string(RoleAdmin)
}
