package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shopping-cart-api-v1/handlers"
)

func SetupRoute(app *fiber.App) {
	router := app.Group("/cart")
	router.Post("/add", handlers.AddCart)
	router.Post("/remove", handlers.RemoveFromCart)
	router.Get("/view", handlers.ViewCart)
	router.Get("/total", handlers.GetTotal)
	router.Delete("/clear", handlers.ClearCart)
}

