package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ApiKeyGuard(validKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		apikey := c.Get("X-API-KEY")

		if apikey == "" || apikey != validKey {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorzed - Invalid API key",
			})
		}
		return c.Next()
	}
}
