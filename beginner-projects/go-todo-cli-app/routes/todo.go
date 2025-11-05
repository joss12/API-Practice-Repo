package routes

import (
	"github.com/go-todo-cli-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
	todo := app.Group("/todos")
	todo.Get("/", handlers.GetTodos)
	todo.Post("/", handlers.CreateTodo)
	todo.Put("/:id", handlers.ToggleTodo)
	todo.Delete("/:id", handlers.DeleteTodo)
}
