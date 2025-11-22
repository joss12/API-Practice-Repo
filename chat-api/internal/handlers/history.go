package handlers

import (
	"context"
	"strconv"

	"github.com/chat-api/internal/repo"
	"github.com/gofiber/fiber/v2"
)

type HistoryHandler struct {
	messages *repo.MessageRepo
}

func NewHistoryHandler(m *repo.MessageRepo) *HistoryHandler {
	return &HistoryHandler{messages: m}
}

func (h *HistoryHandler) GetRoomHistory(c *fiber.Ctx) error {
	room := c.Query("room")
	if room == "" {
		return fiber.NewError(fiber.StatusBadRequest, "room required")
	}

	limitStr := c.Query("limit", "50")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 500 {
		limit = 500
	}

	msgs, err := h.messages.GetLastMessage(context.Background(), room, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "db error")
	}

	//Reverse order: oldest -> neest (UI-Frendfly)
	for i, j := 0, len(msgs)-1; i < j; i, j = i+1, j-1 {
		msgs[i], msgs[j] = msgs[j], msgs[i]
	}
	return c.JSON(msgs)
}
