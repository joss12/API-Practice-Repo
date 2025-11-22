package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/chat-api/internal/config"
	"github.com/gofiber/fiber/v2"
)

func newTestAuthApp(h *AuthHandler) *fiber.App {
	a := fiber.New()
	a.Post("/auth/register", h.Register)
	a.Post("/auth/login", h.Login)
	return a
}

func TestRegisterBadJSON(t *testing.T) {
	h := NewAuthHandler(nil, config.Config{JWTSecret: "x"})
	a := newTestAuthApp(h)

	req := httptest.NewRequest("POST", "/auth/register", bytes.NewBuffer([]byte("{bad json}")))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := a.Test(req)

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}
