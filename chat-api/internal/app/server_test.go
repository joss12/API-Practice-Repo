package app

import (
	"net/http/httptest"
	"testing"

	"github.com/chat-api/internal/handlers"
)

func TestNewCreatesFiberApp(t *testing.T) {
	a := New()
	if a == nil {
		t.Fatal("expected fiber app, got nil")
	}
}

func TestHealthHandler(t *testing.T) {
	a := New()
	a.Get("/health", handlers.Health)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := a.Test(req, -1)
	if err != nil {
		t.Fatalf("fiber test error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, go %d", resp.StatusCode)
	}
}
