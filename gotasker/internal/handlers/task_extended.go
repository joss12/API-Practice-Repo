package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gotasker/internal/database"
	"github.com/gotasker/internal/models"
)

func AssignTask(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "admin only"})
	}

	taskID, _ := strconv.Atoi(c.Params("id"))
	var payload struct {
		UserId uint `json:"user_id"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	var task models.Task
	if err := database.DB.First(&task, taskID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "task not found"})
	}

	task.UserID = payload.UserId
	database.DB.Save(&task)
	return c.JSON(fiber.Map{"message": "task assigned", "task": task})
}

func UpdateStatus(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	role := c.Locals("role").(string)
	taskID, _ := strconv.Atoi(c.Params("id"))

	var task models.Task
	if err := database.DB.First(&task, taskID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "task not found"})
	}

	//Users can only update own tasks unless admin
	if role != "admin" && task.UserID != userID {
		return c.Status(403).JSON(fiber.Map{"error": "not your task"})
	}

	var payload struct {
		Status string `json:"status"`
	}
	c.BodyParser(&payload)

	task.Status = payload.Status
	database.DB.Save(&task)
	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "admin only"})
	}

	taskID, _ := strconv.Atoi(c.Params("id"))
	if err := database.DB.Delete(&models.Task{}, taskID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "delete failed"})
	}
	return c.JSON(fiber.Map{"message": "task delete"})
}

// User: list own tasks
func MyTasks(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var tasks []models.Task
	database.DB.Where("user_id = ?", userID).Find(&tasks)
	return c.JSON(tasks)
}
