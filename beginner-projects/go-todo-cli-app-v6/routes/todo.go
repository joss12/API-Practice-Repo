package routes

import (
	"github.com/go-todo-cli-v6/handlers"
	"github.com/gofiber/fiber/v2"
)

func AppRoute(app *fiber.App) {
	router := app.Group("/todos")
	router.Get("/", handlers.GetTodos)
	router.Post("/", handlers.CreateTodo)
	router.Put("/:id", handlers.ToggleTodos)
	router.Delete("/:id", handlers.DeleteTodo)
}
