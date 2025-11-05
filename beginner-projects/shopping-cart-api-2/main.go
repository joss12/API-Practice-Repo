package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shopping-cart-api-v2/router"
)

func main() {
	app := fiber.New()

	//Logger
	app.Use(logger.New())

	//routes
	router.AppRoute(app)

	app.Listen(":3002")
}
