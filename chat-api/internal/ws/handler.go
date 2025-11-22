package ws

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/chat-api/internal/config"
	"github.com/chat-api/internal/models"
	appredis "github.com/chat-api/internal/redis"
	"github.com/chat-api/internal/repo"
	red "github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type RedisPublisher interface {
	Publish(ctx context.Context, channel string, message interface{}) *redisStatus
}

// We don't need full go-redis type in this file; a minimal interface wrapper
// is enough for testing and production. But in practice, you can just use the
// real *redis.Client type directly if you prefer.
type redisStatus interface {
	Err() error
}

// Routes now receives both DB and Redis.
// db can be nil
// rdb can be nil (fallback)

// func Routes(a *fiber.App, hub *Hub, db *gorm.DB, rdb *red.Client)
func Routes(a *fiber.App, hub *Hub, db *gorm.DB, rdb *red.Client, cfg config.Config) {

	// Optional Message Repo
	var mrepo *repo.MessageRepo
	if db != nil {
		mrepo = repo.NewMessageRepo(db)
	}

	// Allow WebSocket upgrade
	a.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	a.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// 1) JWT from ?token=<jwt>
		tokenStr := c.Query("token")
		if tokenStr == "" {
			_ = c.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "missing token"),
			)
			_ = c.Close()
			return
		}
		claims := jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !tkn.Valid {
			_ = c.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "invalid token"),
			)
			_ = c.Close()
			return
		}

		userIDfloat, ok := claims["userID"].(float64)
		if !ok {
			_ = c.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "invalid userID claim"),
			)
			_ = c.Close()
			return
		}
		userID := uint(userIDfloat)
		//2) Room from ?room=... (default: general)
		room := c.Query("room")
		if room == "" {
			room = "general"
		}

		if rdb != nil {
			_ = appredis.TrackUserJoin(context.Background(), rdb, room, userID)
		}

		client := &Client{
			send:   make(chan []byte, 256),
			room:   room,
			userID: userID,
		}
		hub.register <- client
		defer func() {
			hub.unregister <- client
			if rdb != nil {
				_ = appredis.TrackUserLeave(context.Background(), rdb, room, userID)
			}
		}()

		//writer goroutine
		go func() {
			for msg := range client.send {
				if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
					return
				}
			}
		}()

		//reader loop
		for {
			mt, p, err := c.ReadMessage()
			if err != nil {
				return
			}
			if mt != websocket.TextMessage {
				continue
			}

			//UNMARSHAL INPUT
			var incoming struct {
				Type      string `json:"type"`
				MessageID uint   `json:"messageID"`
				Body      string `json:"body"`
				IsTyping  bool   `json:"isTyping"`
			}
			_ = json.Unmarshal(p, &incoming)

			// ---- HANDLE ACK ----
			if incoming.Type == "ack" && incoming.MessageID > 0 {
				if mrepo != nil {
					_ = mrepo.MarkDelivered(context.Background(), incoming.MessageID, userID)
				}

				// Broadcast delivered event to room
				deliveredEvent, _ := json.Marshal(fiber.Map{
					"type":      "delivered",
					"messageID": incoming.MessageID,
					"userID":    userID,
				})

				b := Broadcast{Room: room, Data: deliveredEvent}
				if rdb != nil {
					_ = rdb.Publish(context.Background(), "chat_message", string(deliveredEvent)).Err()
				} else {
					hub.broadcast <- b
				}
				continue
			}

			// ---- HANDLE NORMAL MESSAGE ----
			if incoming.Type == "message" {
				// store in DB
				var msgID uint = 0
				if mrepo != nil {
					msg := &models.Message{
						UserID:    userID,
						Room:      room,
						Body:      incoming.Body,
						CreatedAt: time.Now(),
					}
					_ = mrepo.Create(context.Background(), msg)
					msgID = msg.ID
				}
				// broadcast event with messageID
				out, _ := json.Marshal(fiber.Map{
					"type":      "message",
					"messageID": msgID,
					"userID":    userID,
					"body":      incoming.Body,
					"room":      room,
				})

				b := Broadcast{Room: room, Data: out}

				if rdb != nil {
					_ = rdb.Publish(context.Background(), "chat_message", string(out)).Err()
				} else {
					hub.broadcast <- b
				}
			}

			//persist message if repo exist
			if mrepo != nil {
				_ = mrepo.Create(context.Background(), &models.Message{
					UserID:    userID,
					Room:      room,
					Body:      string(p),
					CreatedAt: time.Now(),
				})
			}
			b := Broadcast{Room: room, Data: p}

			//publish viia Redis if provided, sels in-process
			if rdb != nil {
				payload, err := json.Marshal(b)
				if err == nil {
					_ = rdb.Publish(context.Background(), "chat_message", string(payload))
				}
			} else {
				hub.broadcast <- b
			}

			if incoming.Type == "seen" && incoming.MessageID > 0 {
				if mrepo != nil {
					_ = mrepo.MarkSeen(context.Background(), incoming.MessageID, userID)
				}

				//Broadcast seen event
				seenEvent, _ := json.Marshal(fiber.Map{
					"type":    "seen",
					"message": incoming.MessageID,
					"userID":  userID,
				})
				b := Broadcast{Room: room, Data: seenEvent}

				if rdb != nil {
					_ = rdb.Publish(context.Background(), "chat_message", string(seenEvent)).Err()
				} else {
					hub.broadcast <- b
				}

				continue
			}

			if incoming.Type == "typing" {
				//Build a typing event to broadcast to others in the room.
				typingEvent, _ := json.Marshal(fiber.Map{
					"type":     "typing",
					"userID":   userID,
					"room":     room,
					"isTyping": incoming.IsTyping,
				})

				b := Broadcast{
					Room: room,
					Data: typingEvent,
				}
				if rdb != nil {
					_ = rdb.Publish(context.Background(), "chat_message", string(typingEvent)).Err()
				} else {
					hub.broadcast <- b
				}
			}
		}
	}))
}
