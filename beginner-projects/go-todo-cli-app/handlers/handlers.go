package handlers

import (
	"strconv"

	"github.com/go-todo-cli-api/models"
	"github.com/gofiber/fiber/v2"
)

var todos = []models.Todo{}
var idCounter = 1

// Getting the todos (Public)
func GetTodos(c *fiber.Ctx) error {
	status := c.Query("completed") //optional filter
	if status == "" {
		return c.JSON(todos)
	}

	var filtered []models.Todo
	for _, t := range todos {
		if (status == "true" && t.Completed) || (status == "falsee" && !t.Completed) {
			filtered = append(filtered, t)
		}
	}
	return c.JSON(filtered)
}

// create todo (private)
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "Cannot parse JSON"})
	}
	todo.ID = idCounter
	idCounter++
	todos = append(todos, *todo)
	return c.Status(201).JSON(todo)
}

// To update todos
func ToggleTodo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = !todos[i].Completed
			return c.JSON(todos[i])
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
			c.SendStatus(204)
			return nil
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
}
