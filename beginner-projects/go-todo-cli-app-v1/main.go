package main

import (
	"github.com/go-todo-cli-api-v1/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	//router files
	routes.Router(app)
	app.Listen(":3002")
}

