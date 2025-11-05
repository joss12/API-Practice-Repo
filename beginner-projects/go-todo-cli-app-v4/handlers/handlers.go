package handlers

import (
	"strconv"

	"github.com/go-todo-cli-api-v4/models"
	"github.com/gofiber/fiber/v2"
)

var todos = []models.Todo{}
var idCouter = 1

// Getting todos
func GetTodos(c *fiber.Ctx) error {
	status := c.Query("completed") // optional filter
	if status == "" {
		return c.JSON(todos)
	}

	var result []models.Todo
	for _, t := range todos {
		if status == "true" && t.Completed || status == "falsse" && !t.Completed {
			result = append(result, t)
		}
	}
	return c.JSON(result)
}

// created todos
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "Cannont pare JSON"})
	}
	todo.ID = idCouter
	idCouter++
	todos = append(todos, *todo)
	return c.Status(201).JSON(fiber.Map{"Message": "Todo created successfully", "todo": todo})
}

// Update todo
func ToggleTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = !todos[i].Completed
			return c.JSON(fiber.Map{"Message": "Todo update successfully", "todo": todos[i]})
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
}

// Deleting todos
func DeleteTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.Set("Message", "Todo deleted successfully")
			c.SendStatus(204)
			return nil
		}
	}
	return c.Status(404).JSON(fiber.Map{"Error": "Todo not found"})
}
