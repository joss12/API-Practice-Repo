package handlers

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/portlog/redis"
	"github.com/portlog/scanner"
)

type ScanRequest struct {
	Host string `json:"host"`
}

type ScanResponse struct {
	Host      string `json:"host"`
	Ports     []int  `json:"open_ports"`
	Cached    bool   `json:"cached"`
	Timestamp string `json:"timestamp"`
}

func ScanHendler(c *fiber.Ctx) error {
	var req ScanRequest

	if err := c.BodyParser(&req); err != nil || req.Host == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or invalid 'host'",
		})
	}

	cacheKey := "scan:" + req.Host

	// Check cache first
	if cached, err := redis.CacheGet(cacheKey); err == nil {
		var ports []int
		if err := json.Unmarshal([]byte(cached), &ports); err == nil {
			return c.JSON(ScanResponse{
				Host:      req.Host,
				Ports:     ports,
				Cached:    true,
				Timestamp: time.Now().Format(time.RFC3339),
			})
		}
	}

	// Not cached, do scan
	ports := scanner.ScanPorts(req.Host)

	// Cache result for 10 minutes
	portBytes, _ := json.Marshal(ports)
	_ = redis.CacheSet(cacheKey, string(portBytes), 10*time.Minute)

	return c.JSON(ScanResponse{
		Host:      req.Host,
		Ports:     ports,
		Cached:    false,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
