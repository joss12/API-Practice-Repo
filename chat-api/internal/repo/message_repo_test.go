package repo

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/chat-api/internal/db"
	"github.com/chat-api/internal/models"
)

func TestMessageRepo_Create_Smoke(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set, skip integration")
	}
	gdb, err := db.Open(dsn)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := gdb.AutoMigrate(&models.Message{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	r := NewMessageRepo(gdb)
	if err := r.Create(context.Background(), &models.Message{
		Body: "hello", CreatedAt: time.Now(),
	}); err != nil {
		t.Fatalf("create: %v", err)
	}
}
