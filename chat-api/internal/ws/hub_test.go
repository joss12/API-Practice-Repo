package ws

import (
	"bytes"
	"testing"
	"time"
)

func TestHubBroadcast(t *testing.T) {
	h := NewHub()
	go h.Run()

	c1 := &Client{send: make(chan []byte, 1), room: "general"}
	c2 := &Client{send: make(chan []byte, 1), room: "general"}
	c3 := &Client{send: make(chan []byte, 1), room: "other"}

	h.register <- c1
	h.register <- c2
	h.register <- c3

	msg := []byte("hello")
	h.broadcast <- Broadcast{Room: "general", Data: msg}

	select {
	case got := <-c1.send:
		if !bytes.Equal(got, msg) {
			t.Fatalf("c1 expected %q, got %q", msg, got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("timeout waiting for c1")
	}

	select {
	case got := <-c2.send:
		if !bytes.Equal(got, msg) {
			t.Fatalf("c2 expected %q, got %q", msg, got)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("timeout waiting for c2")
	}

	select {
	case <-c3.send:
		t.Fatalf("c3 should not receive message for room general")
	case <-time.After(100 * time.Millisecond):
		//OK: no message
	}
}
