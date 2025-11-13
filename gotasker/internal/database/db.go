package database

import (
	"fmt"
	"log"

	"github.com/gotasker/internal/config"
	"github.com/gotasker/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_NAME"),
	)

	//db, err := gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connected to DB:", err)
	}
	DB = db

	db.AutoMigrate(&models.User{}, &models.Task{})
	log.Println("Database connected & migrate")
}
