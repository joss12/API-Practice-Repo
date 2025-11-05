package main

import (
	"github.com/go-todo-cli-v6/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//routes
	routes.AppRoute(app)

	app.Listen(":3002")
}
