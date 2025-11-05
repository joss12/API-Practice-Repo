package routes

import (
	"github.com/go-todo-cli-api-v1/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	rApp := app.Group("/todos")
	rApp.Get("/", handlers.GetTodos)
	rApp.Post("/", handlers.CreateTodo)
	rApp.Put("/:id", handlers.ToggleTodos)
	rApp.Delete("/:id", handlers.DeleteTodo)
}

