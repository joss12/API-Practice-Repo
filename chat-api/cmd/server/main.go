package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/chat-api/internal/app"
	"github.com/chat-api/internal/config"
	"github.com/chat-api/internal/db"
	"github.com/chat-api/internal/handlers"
	"github.com/chat-api/internal/metrics"
	"github.com/chat-api/internal/middleware"
	"github.com/chat-api/internal/models"
	"github.com/chat-api/internal/redis"
	"github.com/chat-api/internal/repo"
	"github.com/chat-api/internal/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	a := app.New()

	// WebSocket Hub
	hub := ws.NewHub()
	go hub.Run()

	// Config
	cfg := config.Load()

	// Database
	gdb, err := db.Open(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := gdb.AutoMigrate(&models.User{}, &models.Message{}); err != nil {
		log.Fatal(err)
	}

	// Redis
	rdb := redis.NewRedis("redis:6379", "")
	sub := redis.Subscribe(rdb, "chat_message")
	metrics.RegisterMetrics()

	// Redis Listener goroutine
	go func() {
		for msg := range sub.Channel() {
			var b ws.Broadcast
			if err := json.Unmarshal([]byte(msg.Payload), &b); err == nil {
				hub.RedisBroadcast(b)
			}
		}
	}()
	// User repo & Auth
	userRepo := repo.NewUserRepo(gdb)
	auth := handlers.NewAuthHandler(userRepo, cfg)

	msgRepo := repo.NewMessageRepo(gdb)
	history := handlers.NewHistoryHandler(msgRepo)

	rooms := handlers.NewRoomHandler(rdb)
	roomHandler := handlers.NewRoomHandler(rdb)

	metrics.RegisterMetrics()
	admin := handlers.NewAdminHandler(rdb)

	// Auth routes
	a.Post("/auth/register", auth.Register)
	a.Post("/auth/login", auth.Login)
	a.Get("/auth/me", middleware.JWTAuth(cfg), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"userID": c.Locals("userID"),
		})
	})

	a.Get("/chat/history", middleware.JWTAuth(cfg), history.GetRoomHistory)
	a.Get("/rooms/presence", middleware.JWTAuth(cfg), rooms.Presence)

	a.Post("/room", middleware.JWTAuth(cfg), roomHandler.CreateRoom)
	a.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	adminGroup := a.Group("/admin", middleware.JWTAuth(cfg), middleware.AdminOnly())

	adminGroup.Get("/rooms", admin.ListRooms)
	adminGroup.Get("/users", admin.ListUsers)
	adminGroup.Post("/ban/:id", admin.BanUser)
	adminGroup.Post("/mute/:id", admin.MuteUser)
	adminGroup.Post("/unmute/:id", admin.UnmuteUser)

	// Health route
	a.Get("/health", handlers.Health)

	// WebSocket routes (FIXED HERE)
	ws.Routes(a, hub, gdb, rdb, cfg)

	// Start server
	addr := ":" + env("PORT", "8080")
	if err := a.Listen(addr); err != nil {
		log.Fatal(err)
	}
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
