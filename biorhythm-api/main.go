package main

import (
	"log"
	"net/http"

	"github.com/biorhythm-api/config"
	"github.com/biorhythm-api/handlers"
	"github.com/biorhythm-api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router with custom config
	router := gin.New()

	// Apply middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Custom error handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Endpoint not found",
			"error":   "The requested endpoint does not exist",
		})
	})

	// Initialize handlers
	biorhythmHandler := handlers.NewBiorhythmHandler()
	healthHandler := handlers.NewHealthHandler()

	// Root welcome route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Welcome to Biorhythm Calculator API",
			"version": "1.0.0",
			"endpoints": gin.H{
				"health":        "GET /api/v1/health",
				"biorhythm":     "GET /api/v1/biorhythm/:birthdate",
				"specific_date": "GET /api/v1/biorhythm/:birthdate/:date",
				"critical_days": "GET /api/v1/critical-days/:birthdate/:days",
				"compatibility": "GET /api/v1/compatibility/:birthdate1/:birthdate2",
				"chart":         "GET /api/v1/chart/:birthdate/:start/:end",
			},
			"example": "GET /api/v1/biorhythm/1990-05-15",
		})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", healthHandler.HealthCheck)

		// Biorhythm endpoints - FIXED: Remove 'error' return type
		api.GET("/biorhythm/:birthdate", func(c *gin.Context) {
			biorhythmHandler.GetBiorhythm(c)
		})
		api.GET("/biorhythm/:birthdate/:date", func(c *gin.Context) {
			biorhythmHandler.GetBiorhythmForDate(c)
		})
		api.GET("/critical-days/:birthdate/:days", func(c *gin.Context) {
			biorhythmHandler.GetCriticalDays(c)
		})
		api.GET("/compatibility/:birthdate1/:birthdate2", func(c *gin.Context) {
			biorhythmHandler.GetCompatibility(c)
		})
		api.GET("/chart/:birthdate/:start/:end", func(c *gin.Context) {
			biorhythmHandler.GetBiorhythmChart(c)
		})
	}

	// Start server
	log.Printf("🧬 Biorhythm API server starting on port %s", cfg.Port)
	log.Printf("📍 API Base URL: http://localhost:%s/api/v1", cfg.Port)
	log.Printf("📋 API Documentation: http://localhost:%s", cfg.Port)
	log.Fatal(router.Run(":" + cfg.Port))
}
