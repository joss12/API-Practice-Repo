package ws

import (
	"encoding/json"
	"net"
	"net/url"
	"testing"

	"github.com/chat-api/internal/config"
	//	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
)

func startAppWSSeen() (*fiber.App, string, func()) {
	hub := NewHub()
	go hub.Run()

	cfg := config.Config{JWTSecret: "testsecret"}
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	Routes(app, hub, nil, nil, cfg)

	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	go app.Listener(ln)

	return app, "ws://" + ln.Addr().String() + "/ws", func() { _ = ln.Close() }
}

func TestWSSeen(t *testing.T) {
	_, base, shutdown := startAppWSSeen()
	defer shutdown()

	u, _ := url.Parse(base)
	q := u.Query()
	q.Set("token", generateToken("testsecret"))
	q.Set("room", "seenroom")
	u.RawQuery = q.Encode()

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()

	seen, _ := json.Marshal(map[string]any{
		"type":      "seen",
		"messageID": 77,
	})

	err = c.WriteMessage(websocket.TextMessage, seen)
	if err != nil {
		t.Fatalf("write: %v", err)
	}
}
