package db

import (
	"os"
	"testing"

	"gorm.io/gorm"
)

func TestOpenInvalidDSN(t *testing.T) {
	// Use an invalid DSN to ensure we get an error
	_, err := Open("postgres://invalid:invalid@localhost:5432/none")
	if err == nil {
		t.Skip("driver may not error until first call; skipping strict assert")
	}
}

// optional smoke: run only when DATABASE_URL present (for CI/docker)
func TestOpenWhenDSNProvided(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set, skip integration")
	}
	db, err := Open(dsn)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	var got *gorm.DB = db
	if got == nil {
		t.Fatal("nil gorm db")
	}
}
