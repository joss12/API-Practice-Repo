package main

import (
	"github.com/go-todo-cli-v5/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//routes
	routes.AppRoute(app)

	app.Listen(":3002")

}
