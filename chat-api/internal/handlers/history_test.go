package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/chat-api/internal/repo"
	"github.com/gofiber/fiber/v2"
)

func TestHistoryMissingRoom(t *testing.T) {
	//Don't need a real repo for the "missing room" test.
	app := fiber.New()

	// We MUST mount the handler on the correct route
	h := &HistoryHandler{messages: &repo.MessageRepo{}} // fake repo OK
	app.Get("/chat/history", h.GetRoomHistory)

	// Missing ?room= param
	req := httptest.NewRequest("GET", "/chat/history", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("fiber test error: %v", err)
	}

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}

	// h := &HistoryHandler{messages: nil}
	// app := fiber.New()
	// app.Get("/chat/history", h.GetRoomHistory)
	//
	// req := httptest.NewRequest("GET", "/chat./history", nil)
	// resp, _ := app.Test(req)
	//
	//	if resp.StatusCode != fiber.StatusBadRequest {
	//		t.Fatalf("expected 400, got %d", resp.StatusCode)
	//	}
}
