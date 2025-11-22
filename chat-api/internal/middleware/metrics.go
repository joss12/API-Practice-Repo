package middleware

import (
	"strconv"

	"github.com/chat-api/internal/metrics"
	"github.com/gofiber/fiber/v2"
)

func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		status := c.Response().StatusCode()
		statusStr := strconv.Itoa(status)

		metrics.RequestCounter.WithLabelValues(
			c.Method(),
			c.Path(),
			statusStr,
		).Inc()

		return err
	}
}
