package ws

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWSTypingBroadcast(t *testing.T) {
	_, base, shutdown := startAppWS()
	defer shutdown()

	u, _ := url.Parse(base)
	q := u.Query()
	q.Set("token", generateToken("testsecret")) // MUST MATCH startAppWS()
	q.Set("room", "typingroom")
	u.RawQuery = q.Encode()

	c1, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("dial c1: %v", err)
	}
	defer c1.Close()

	c2, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("dial c2: %v", err)
	}
	defer c2.Close()

	// typing event
	payload, _ := json.Marshal(map[string]interface{}{
		"type":     "typing",
		"isTyping": true,
	})

	if err := c1.WriteMessage(websocket.TextMessage, payload); err != nil {
		t.Fatalf("write typing: %v", err)
	}

	_, msg, err := c2.ReadMessage()
	if err != nil {
		t.Fatalf("read typing: %v", err)
	}

	var event map[string]interface{}
	_ = json.Unmarshal(msg, &event)

	if event["type"] != "typing" {
		t.Fatalf("expected type typing, got %v", event["type"])
	}
}
