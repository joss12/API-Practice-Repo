package main

import (
	"github.com/go-todo-cli-api-v2/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	//Routes
	routes.AppRoute(app)
	app.Listen(":3002")
}
