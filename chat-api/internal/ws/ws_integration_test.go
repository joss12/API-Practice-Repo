package ws

import (
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"

	"github.com/chat-api/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// Helper: generate JWT for tests

func generateToken(secret string) string {
	claims := jwt.MapClaims{
		"userID": 1,
		"exp":    time.Now().Add(time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, _ := t.SignedString([]byte(secret))
	return s
}

// Start Fiber on a real TCP listener
func startFiberApp(app *fiber.App) (string, func()) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	go func() {
		if err := app.Listener(ln); err != nil {
			panic(err)
		}
	}()
	return "http://" + ln.Addr().String(), func() { _ = ln.Close() }
}

func TestWebSocketRoomBroadcast(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	cfg := config.Config{JWTSecret: "testsecret"}

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	// IMPORTANT: use fiberws (NOT gorilla's websocket) for server routes
	Routes(app, hub, nil, nil, cfg)

	baseURL, shutdown := startFiberApp(app)
	defer shutdown()

	// Build ws:// URL with room + token
	u, _ := url.Parse(baseURL)
	u.Scheme = "ws"
	u.Path = "/ws"
	q := u.Query()
	q.Set("room", "testroom")
	q.Set("token", generateToken("testsecret"))
	u.RawQuery = q.Encode()

	// Connect client 1
	c1, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("dial c1 failed: %v", err)
	}
	defer c1.Close()

	// Connect client 2
	c2, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("dial c2 failed: %v", err)
	}
	defer c2.Close()

	// Send message from c1
	want := "hello"
	if err := c1.WriteMessage(websocket.TextMessage, []byte(want)); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	// c2 should receive it
	_, msg, err := c2.ReadMessage()
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}
	if string(msg) != want {
		t.Fatalf("expected %q, got %q", want, string(msg))
	}
}
