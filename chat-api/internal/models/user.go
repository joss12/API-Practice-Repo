package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"index"`
	Body         string    `gorm:"type:text;not null"`
	Username     string    `gorm:"uniqueIndex;size:64;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
