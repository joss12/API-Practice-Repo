package controllers

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"time"

	"testing"

	"github.com/auth-system/config"
	"github.com/auth-system/database"
	"github.com/auth-system/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setup test app with in-memory sql instead of Mysql
func setupTestApp() *fiber.App {
	app := fiber.New()

	//use memero DB for test -> isolated and faste
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.DB = db
	database.DB.AutoMigrate(&models.User{})

	config.AppConfig = &config.Config{
		JWTSecret: "test-secret",
		Port:      "9999",
	}

	//Routes
	app.Post("/register", Register)
	app.Post("/login", Login)
	app.Post("/forgot-password", ForgotPassword)
	app.Post("/reset-password", ResetPassword)

	return app
}

func TestRegister(t *testing.T) {
	app := setupTestApp()

	body := map[string]string{
		"name":     "Eddy",
		"email":    "eddy@example.com",
		"password": "mypassword",
	}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(payload))
	req.Header.Set("content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

}

func TestLogin(t *testing.T) {
	app := setupTestApp()

	//Hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte("secret"), 12)

	//Insert user into test DB (use shared db)
	user := models.User{
		Name:     "Tester",
		Email:    "login@test.com",
		Password: string(hashed),
		Role:     models.RoleUser,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	database.DB.Create(&user)

	//Attemp login
	body := map[string]string{
		"email":    "login@test.com",
		"password": "secret",
	}
	data, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Test request failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}
}

func TestForgotPassword(t *testing.T) {
	app := setupTestApp()

	//Create user
	user := models.User{
		Name:  "Forget",
		Email: "forgot@test.com",
	}
	database.DB.Create(&user)

	body := map[string]string{"email": "forgot@test.com"}
	data, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}
}

func TestResetPassword(t *testing.T) {
	app := setupTestApp()

	//Create user with reset token
	token := "resettoken123"
	expiry := time.Now().Add(15 * time.Minute)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("oldpass"), 12)

	user := models.User{
		Name:             "Reset",
		Email:            "reset@test.com",
		Password:         string(hashed),
		ResetToken:       token,
		ResetTokenExpiry: &expiry,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create reset user: %v", err)
	}

	//database.DB.Create(&user)

	body := map[string]string{
		"token":        token,
		"new_password": "newpassword",
	}
	data, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/reset-password", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Reset test failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}

}
