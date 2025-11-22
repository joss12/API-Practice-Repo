package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRoomsPresenceUnavailable(t *testing.T) {
	h := &RoomHandler{rdb: nil}
	app := fiber.New()
	app.Get("/rooms/presence", h.Presence)

	req := httptest.NewRequest("GET", "/rooms/presence", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("fiber test error: %v", err)
	}

	if resp.StatusCode != fiber.StatusServiceUnavailable {
		t.Fatalf("expected 503, got %d", resp.StatusCode)
	}
}
