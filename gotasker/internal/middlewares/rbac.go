package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotasker/internal/database"
	"github.com/gotasker/internal/models"
)

func RequiresRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("user").(map[string]interface{})["id"]

		var user models.User
		database.DB.First(&user, userId)

		if user.Role != role {
			return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
		}
		return c.Next()
	}
}
