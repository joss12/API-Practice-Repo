package router

import (
	"github.com/blog-api-v3/handlers"
	"github.com/gofiber/fiber/v2"
)

func AppRouter(app *fiber.App) {
	router := app.Group("/posts")

	router.Post("/create", handlers.CreatePost)
	router.Get("/", handlers.GetAllPosts)
	router.Get("/:id", handlers.GetPostByID)
	router.Put("/:id", handlers.UpdatePost)
	router.Delete("/:id", handlers.DeletePost)
}
