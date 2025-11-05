package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL     string
	JWTSecret string
	Port      string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using system env variable")
	}

	AppConfig = &Config{
		DBURL:     getEnv("DB_URL", "root:password@tcp(localhost:3306)/authdb?parseTime=true"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		Port:      getEnv("PORT", "3003"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
