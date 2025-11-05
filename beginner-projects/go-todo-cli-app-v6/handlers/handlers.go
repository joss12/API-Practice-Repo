package handlers

import (
	"strconv"

	"github.com/go-todo-cli-v6/models"
	"github.com/gofiber/fiber/v2"
)

var todos = []models.Todo{}
var idCounter = 1

// Getting todos
func GetTodos(c *fiber.Ctx) error {
	status := c.Query("completed") //optional filte
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

// create todos
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "Cannot parse JSON"})
	}
	todo.ID = idCounter
	idCounter++
	todos = append(todos, *todo)
	return c.Status(201).JSON(fiber.Map{"Message": "Todo create successfully", "todo": todo})
}

// To Update todos by making it completed
func ToggleTodos(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i, t := range todos {
		if t.ID == id {
			todos[i].Completed = !todos[i].Completed
			return c.JSON(fiber.Map{"Message": "Todo updated successfully", "todo": todos[i]})
		}
	}
	return c.Status(404).JSON(fiber.Map{"Error": "Todo no found"})
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
	return c.Status(404).JSON(fiber.Map{"Error": "Todo not fond"})

}
