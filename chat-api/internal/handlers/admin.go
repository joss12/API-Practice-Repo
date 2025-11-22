package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type AdminHandler struct {
	Rdb *redis.Client
}

func NewAdminHandler(rdb *redis.Client) *AdminHandler {
	return &AdminHandler{Rdb: rdb}
}

// ---- LIST ROOMS ----
func (h *AdminHandler) ListRooms(c *fiber.Ctx) error {
	if h.Rdb == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "redis unavailable")
	}

	rooms, err := h.Rdb.SMembers(context.Background(), "rooms:active").Result()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "redis error")
	}

	type RoomInfo struct {
		Name  string `json:"name"`
		Users int64  `json:"users"`
	}

	out := []RoomInfo{}
	for _, room := range rooms {
		count, _ := h.Rdb.SCard(context.Background(), "room:"+room+":online").Result()
		out = append(out, RoomInfo{Name: room, Users: count})
	}

	return c.JSON(out)
}

// ---- LIST ONLINE USERS ----
func (h *AdminHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.Rdb.SMembers(context.Background(), "users:online").Result()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "redis error")
	}
	return c.JSON(users)
}

// ---- BAN USER ----
func (h *AdminHandler) BanUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.Rdb.SAdd(context.Background(), "users:banned", id).Err()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "redis error")
	}
	return c.JSON(fiber.Map{"banned": id})
}

// ---- MUTE USER ----
func (h *AdminHandler) MuteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.Rdb.SAdd(context.Background(), "users:muted", id).Err()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "redis error")
	}

	return c.JSON(fiber.Map{"muted": id})
}

// ---- UNMUTE USER ----
func (h *AdminHandler) UnmuteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.Rdb.SRem(context.Background(), "users:muted", id).Err()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "redis error")
	}
	return c.JSON(fiber.Map{"unmuted": id})
}
