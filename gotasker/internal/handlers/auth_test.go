package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gotasker/internal/config"
	"github.com/gotasker/internal/database"
)

func setupTestApp() *fiber.App {
	config.LoadEnv()
	database.ConnectDB()
	app := fiber.New()
	app.Post("/register", Register)
	app.Post("/login", Login)
	return app
}

func TestRegisterAndLogin(t *testing.T) {
	app := setupTestApp()

	//register
	body := map[string]string{
		"name":     "Tester",
		"email":    "tester@email.com",
		"password": "1234",
		"role":     "user",
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	if resp.StatusCode != 200 {
		t.Fatalf("register failed, status code: %d", resp.StatusCode)
	}

	//login
	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = app.Test(req)
	if resp.StatusCode != 200 {
		t.Fatalf("login failed: %v", resp.StatusCode)
	}

}
