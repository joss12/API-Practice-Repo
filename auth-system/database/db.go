package database

import (
	"log"

	"github.com/auth-system/config"
	"github.com/auth-system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {

	var err error
	dsn := config.AppConfig.DBURL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("Connected to MySQL")

	// Drop existing table and recreate (ONLY for development!)
	DB.Migrator().DropTable(&models.User{})

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate User model: %v", err)
	}
	log.Println("Auto-migrated User model.")

}
