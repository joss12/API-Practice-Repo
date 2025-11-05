package handlers

import (
	"strconv"

	"github.com/go-todo-cli-v5/models"
	"github.com/gofiber/fiber/v2"
)

var todos = []models.Todo{}
var idCounter = 1

func GetTodos(c *fiber.Ctx) error {
	status := c.Query("completed")
	if status == "" {
		return c.JSON(todos)
	}

	var result []models.Todo
	for _, t := range todos {
		if status == "true" && t.Completed || status == "false" && !t.Completed {
			result = append(result, t)
		}
	}
	return c.JSON(result)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(404).JSON(fiber.Map{"Error": "Cannont parse JSON"})
	}
	todo.ID = idCounter
	idCounter++
	todos = append(todos, *todo)
	return c.Status(201).JSON(fiber.Map{"Message": "Todo created successfully", "todo": todo})
}

// To update todos by making it completed
func UpdateTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = !todos[i].Completed
			return c.JSON(fiber.Map{"Message": "Todo updated successfully", "todo": todos[i]})
		}
	}
	return c.Status(404).JSON(fiber.Map{"Error": "Todo not found"})
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
	return c.Status(404).JSON(fiber.Map{"Message": "Todo not found"})
}
