package routes

import (
	"github.com/auth-system/controllers"
	"github.com/auth-system/database"
	"github.com/auth-system/middleware"
	"github.com/auth-system/models"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	router := app.Group("/api")

	router.Post("/register", controllers.Register)
	router.Post("/login", controllers.Login)
	router.Post("/forgot-password", controllers.ForgotPassword)
	router.Post("/reset-password", controllers.ResetPassword)

	//Protected routes
	auth := router.Group("/auth", middleware.RequireAuth)
	auth.Get("/me", func(c *fiber.Ctx) error {
		user := c.Locals("user").(models.User)
		return c.JSON(fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		})
	})

	//Admin-only route
	admin := auth.Group("/admin", middleware.RequireRole(models.RoleAdmin))

	admin.Get("/users", controllers.GetAllUsers)
	admin.Put("/promote/:id", controllers.PromoteUser)
	admin.Delete("/delete/:id", controllers.DeleteUser)

	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome Admin"})
	})

	router.Put("/promote/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user models.User
		if err := database.DB.First(&user, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		user.Role = models.RoleAdmin
		database.DB.Save(&user)
		return c.JSON(fiber.Map{"error": "User promoted to admin"})
	})

}
