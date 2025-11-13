package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotasker/internal/config"
	"github.com/gotasker/internal/database"
	"github.com/gotasker/internal/router"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	app := fiber.New()
	router.Setup(app)

	app.Listen(":" + config.GetEnv("PORT"))
}
