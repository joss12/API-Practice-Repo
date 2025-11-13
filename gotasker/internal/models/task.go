package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	Status string `json:"status"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
}
