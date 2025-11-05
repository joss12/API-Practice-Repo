package main

import (
	"github.com/go-todo-cli-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetupTodoRoutes(app)
	app.Listen(":3002")
}
