package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestCreateRoomBadJSON(t *testing.T) {
	app := fiber.New()

	h := &RoomHandler{rdb: nil}
	app.Post("/rooms", h.CreateRoom)

	req := httptest.NewRequest("POST", "/rooms", bytes.NewBuffer([]byte("{bad")))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestCreateRoomOK(t *testing.T) {
	app := fiber.New()

	h := &RoomHandler{rdb: nil}
	app.Post("/rooms", h.CreateRoom)

	body := []byte(`{"name": "myroom"}`)
	req := httptest.NewRequest("POST", "/rooms", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
}
