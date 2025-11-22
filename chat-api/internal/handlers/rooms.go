package handlers

import (
	"context"

	appredis "github.com/chat-api/internal/redis"
	"github.com/gofiber/fiber/v2"
	red "github.com/redis/go-redis/v9"
)

type RoomHandler struct {
	rdb *red.Client
}

func NewRoomHandler(rdb *red.Client) *RoomHandler {
	return &RoomHandler{rdb: rdb}
}

type RoomInfo struct {
	Name      string `json:"name"`
	UserCount int64  `json:"userCount"`
}

type createRoomReq struct {
	Name string `json:"name"`
}

// GET /rooms/presence
// Returns total online users and per-room counts.
func (h *RoomHandler) Presence(c *fiber.Ctx) error {
	if h.rdb == nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "presence unavailable")
	}

	ctx := context.Background()

	online, err := appredis.GetOnlineUsers(ctx, h.rdb)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "redis error")
	}

	roomNames, err := appredis.GetRooms(ctx, h.rdb)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "redis error")
	}

	infos := make([]RoomInfo, 0, len(roomNames))
	for _, r := range roomNames {
		cnt, err := appredis.GetRoomUserCount(ctx, h.rdb, r)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "redis error")
		}
		infos = append(infos, RoomInfo{
			Name:      r,
			UserCount: cnt,
		})
	}

	return c.JSON(fiber.Map{
		"totalOnline": len(online),
		"rooms":       infos,
	})
}

func (h *RoomHandler) CreateRoom(c *fiber.Ctx) error {

	var body createRoomReq
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON")
	}

	if len(body.Name) < 2 {
		return fiber.NewError(fiber.StatusBadRequest, "room name too short")
	}

	//Save to redis
	if h.rdb != nil {
		err := h.rdb.SAdd(context.Background(), "rooms:active", body.Name).Err()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "redis error")
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ok":   true,
		"room": body.Name,
	})
}
