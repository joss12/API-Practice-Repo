package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	gorm.Model
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	Password         string     `json:"-"`
	Role             Role       `json:"role" gorm:"default:user"`
	ResetToken       string     `json:"-" gorm:"column:reset_token"`
	ResetTokenExpiry *time.Time `json:"-" gorm:"column:reset_token_expiry"`
}
