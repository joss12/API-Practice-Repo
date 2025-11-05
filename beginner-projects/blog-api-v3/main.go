package main

import (
	"github.com/blog-api-v3/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	//Logger
	app.Use(logger.New())
	router.AppRouter(app)

	app.Listen(":3002")
}
