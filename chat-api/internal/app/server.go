package app

import "github.com/gofiber/fiber/v2"

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		//sensible defaults; no banners, strict routing off
		AppName:               "chatAPI",
		DisableStartupMessage: true,
	})
	return app
}
