package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shopping-cart-api/handlers"
)

func SetupCartRoutes(app *fiber.App) {
	router := app.Group("/cart")

	router.Post("/add", handlers.AddToCart)
	router.Post("/remove", handlers.RemoveFromCart)
	router.Get("/view", handlers.ViewCart)
	router.Get("/total", handlers.ClearCart)
	router.Delete("/clear", handlers.ClearCart)
}
