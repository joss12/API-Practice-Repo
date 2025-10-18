package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/portlog/config"
	"github.com/portlog/handlers"
	"github.com/portlog/middleware"
	"github.com/portlog/redis"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.LoadConfig()

	//Logger setup
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("🚀 Starting PortLog API...")

	//Connect to Redis
	redis.Connect(cfg.RedisAddr, cfg.RedisPass)

	app := fiber.New()

	//Middleware: APII key check
	app.Use(middleware.ApiKeyGuard(cfg.ApiKey))

	//Route registraton
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Post("/scan", handlers.ScanHendler)

	//Start  server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Info().Msgf("Listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("Server crashed!")
	}
}
