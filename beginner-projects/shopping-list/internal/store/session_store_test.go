package store

import (
	"testing"
	"time"

	"github.com/shopping-list-backend/internal/models"
)

func TestSessionStore_Basic(t *testing.T) {
	store := NewSessionStore()

	// Create session
	session, _ := models.NewSession("john", models.RoleUser)
	store.Create(session)

	// Get session
	retrieved, err := store.Get(session.ID)
	if err != nil || retrieved.Username != "john" {
		t.Error("Should retrieve session")
	}

	// Delete session
	store.Delete(session.ID)
	_, err = store.Get(session.ID)
	if err != models.ErrInvalidSession {
		t.Error("Deleted session should not exist")
	}
}

func TestSessionStore_Expired(t *testing.T) {
	store := NewSessionStore()

	// Create and expire session
	session, _ := models.NewSession("john", models.RoleUser)
	session.ExpiresAt = time.Now().Add(-1 * time.Hour)
	store.Create(session)

	// Should return expired error
	_, err := store.Get(session.ID)
	if err != models.ErrSessionExpired {
		t.Error("Should return expired error")
	}

	// Cleanup should remove it
	count := store.CleanupExpired()
	if count != 1 {
		t.Errorf("Should cleanup 1 session, got %d", count)
	}
}
