package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("DB_URL", "test-db")
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("PORT", "9999")

	LoadConfig()

	if AppConfig.DBURL != "test-db" {
		t.Errorf("Expected DB_URL to be 'test-db', go %s", AppConfig.DBURL)
	}

	if AppConfig.JWTSecret != "test-secret" {
		t.Errorf("Expected JWT_SECRET to be 'test-secret', go %s", AppConfig.JWTSecret)
	}

	if AppConfig.Port != "9999" {
		t.Errorf("Expected PORT to be '9999', got %s", AppConfig.Port)
	}
}
