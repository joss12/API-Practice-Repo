package models

import "time"

type Message struct {
	ID          uint       `gorm:"primaryKey"`
	UserID      uint       `gorm:"index"`
	Room        string     `gorm:"size:64;index;not null;default:'general'"`
	Body        string     `gorm:"type:text;not null"`
	CreatedAt   time.Time  `grom:"autoCreateTime"`
	DeliveredAt *time.Time `json:"deliveredAt"`
	SeenAt      *time.Time `json:"seenAt"`
}
