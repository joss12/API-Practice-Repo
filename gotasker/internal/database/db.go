package database

import (
	"fmt"
	"log"

	"github.com/gotasker/internal/confing"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		confing.GetEnv("DB_USER"),
		confing.GetEnv("DB_PASSWORD"),
		confing.GetEnv("DB_HOST"),
		confing.GetEnv("DB_PORT"),
		confing.GetEnv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connected to DB:", err)
	}
	DB = db

	db.AutoMigrate(&models.User{}, &models.Task{})
	log.Println("Database connected & migrate")
}
