package middleware

import "github.com/gofiber/fiber/v2"

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		if role != "admin" {
			return fiber.ErrForbidden
		}
		return c.Next()
	}
}
