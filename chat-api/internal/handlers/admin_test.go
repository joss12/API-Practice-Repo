package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestAdminListRooms_NoRedis(t *testing.T) {
	app := fiber.New()
	h := &AdminHandler{Rdb: nil}

	app.Get("/admin/rooms", h.ListRooms)

	req := httptest.NewRequest("GET", "/admin/rooms", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", resp.StatusCode)
	}
}
