package middleware

import (
	"github.com/chat-api/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("Authorization")
		if tokenStr == "" {
			return fiber.ErrUnauthorized
		}

		claims := jwt.MapClaims{}
		t, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !t.Valid {
			return fiber.ErrUnauthorized
		}

		userID := uint(claims["userID"].(float64))
		c.Locals("userID", userID)
		return c.Next()
	}
}
