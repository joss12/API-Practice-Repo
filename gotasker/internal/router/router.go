package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotasker/internal/handlers"
	"github.com/gotasker/internal/middlewares"
)

func Setup(app *fiber.App) {

	api := app.Group("/api/v1")

	// Auth routes
	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)
	api.Patch("/users/promote/:id", middlewares.JWTMiddleware(), handlers.PromoteUser)
	api.Post("/logout", middlewares.JWTMiddleware(), handlers.Logout)

	// Protected group
	tasks := api.Group("/tasks")
	tasks.Use(middlewares.JWTMiddleware())

	// Task routes
	tasks.Post("/", handlers.CreateTasks)
	tasks.Get("/", handlers.ListTasks)
	tasks.Get("/my", handlers.MyTasks)
	tasks.Get("/admin", handlers.AdminListTasks)
	tasks.Patch("/assign/:id", handlers.AssignTask)
	tasks.Patch("/status/:id", handlers.UpdateStatus)
	tasks.Delete("/:id", handlers.DeleteTask)

}
