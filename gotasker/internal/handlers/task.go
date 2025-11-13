package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotasker/internal/database"
	"github.com/gotasker/internal/models"
)

func CreateTasks(c *fiber.Ctx) error {
	//var task models.Task
	//c.BodyParser(&task)
	//database.DB.Create(&task)
	//return c.JSON(task)

	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	userID := c.Locals("user_id").(uint)
	task.UserID = userID

	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(task)
}

func ListTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	database.DB.Find(&tasks)
	return c.JSON(tasks)
}
