package routes

import (
	"github.com/go-todo-cli-api-v4/handlers"
	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App) {
	router := app.Group("/todos")

	router.Get("/", handlers.GetTodos)
	router.Post("/", handlers.CreateTodo)
	router.Put("/:id", handlers.ToggleTodo)
	router.Delete("/:id", handlers.DeleteTodo)
}
