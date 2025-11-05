package main

import (
	"github.com/blog-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//routes
	routes.AppRoute(app)

	app.Listen(":3002")
}

