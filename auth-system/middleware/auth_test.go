package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/auth-system/config"
	"github.com/auth-system/database"
	"github.com/auth-system/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupMiddlewareTestApp() *fiber.App {
	app := fiber.New()

	// In-memory SQLite DB (isolated for testing)
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.DB = db
	database.DB.AutoMigrate(&models.User{})

	// Insert test user
	user := models.User{
		Name:     "Admin",
		Email:    "admin@example.com",
		Password: "testpass",
		Role:     models.RoleAdmin,
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	// Make sure JWT secret is known in test
	config.AppConfig = &config.Config{
		JWTSecret: "test-secret",
	}

	// Protected route for test
	app.Get("/admin", RequireAuth, RequireRole(models.RoleAdmin), func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return app
}

func TestRequireAuthAndRole(t *testing.T) {
	app := setupMiddlewareTestApp()

	// Retrieve user ID directly from DB (could be 1, could be something else)
	var user models.User
	database.DB.First(&user)

	// Create JWT using the correct ID and role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": string(user.Role),
		"exp":  time.Now().Add(time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+signedToken)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
}
