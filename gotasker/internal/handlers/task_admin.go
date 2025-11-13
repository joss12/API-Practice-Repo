package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotasker/internal/database"
	"github.com/gotasker/internal/models"
)

func AdminListTasks(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "admin only"})
	}

	var tasks []models.Task
	if err := database.DB.Preload("User").Find(&tasks).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(tasks)
}

func PromoteUser(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "admin only"})
	}

	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	user.Role = "admin"
	database.DB.Save(&user)
	return c.JSON(fiber.Map{"message": "user promoted", "user": user})
}
