package main

import (
	"log"

	"github.com/auth-system/config"
	"github.com/auth-system/database"
	"github.com/auth-system/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("> Starting App...")

	config.LoadConfig()
	log.Printf("> Loaded Config: DB_URL=%s", config.AppConfig.DBURL)

	database.ConnectDB()
	log.Println("> Connected to DB")

	app := fiber.New()
	routes.Setup(app)

	log.Printf("> Server running on port %s", config.AppConfig.Port)
	log.Fatal(app.Listen(":" + config.AppConfig.Port))
}
