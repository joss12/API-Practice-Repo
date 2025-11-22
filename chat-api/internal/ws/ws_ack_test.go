package ws

import (
	"encoding/json"
	"net"
	"net/url"
	"testing"

	"github.com/chat-api/internal/config"
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
)

func startAppWS() (*fiber.App, string, func()) {
	hub := NewHub()
	go hub.Run()

	cfg := config.Config{JWTSecret: "testsecret"}
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	Routes(app, hub, nil, nil, cfg)

	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	go app.Listener(ln)

	return app, "ws://" + ln.Addr().String() + "/ws", func() { _ = ln.Close() }
}

func TestWSAck(t *testing.T) {
	_, base, shutdown := startAppWS()
	defer shutdown()

	u, _ := url.Parse(base)
	q := u.Query()

	// Use the existing test token helper from ws_integration_test.go
	q.Set("token", generateToken("testsecret"))
	q.Set("room", "ackroom")

	u.RawQuery = q.Encode()

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()

	ack, _ := json.Marshal(map[string]interface{}{
		"type":      "ack",
		"messageID": 99,
	})

	err = c.WriteMessage(websocket.TextMessage, ack)
	if err != nil {
		t.Fatalf("write: %v", err)
	}
}
