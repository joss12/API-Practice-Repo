package models

import (
	"testing"
	"time"
)

func TestNewSession(t *testing.T) {
	tests := []struct {
		name     string
		username string
		role     Role
	}{
		{"admin session", "admin", RoleAdmin},
		{"user session", "john", RoleUser},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session, err := NewSession(tt.username, tt.role)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if session.Username != tt.username {
				t.Errorf("expected username %s, got %s", tt.username, session.Username)
			}

			if session.Role != tt.role {
				t.Errorf("expected role %s, got %s", tt.role, session.Role)
			}

			if session.ID == "" {
				t.Errorf("session ID should not be empty")
			}

			// Session ID should be 64 chars (32 bytes hex encoded)
			if len(session.ID) != 64 {
				t.Errorf("expected session ID length 64, got %d", len(session.ID))
			}

			if session.CreatedAt.IsZero() {
				t.Errorf("CreatedAt should be set")
			}

			if session.ExpiresAt.IsZero() {
				t.Errorf("ExpiresAt should be set")
			}

			// ExpiresAt should be after CreatedAt
			if !session.ExpiresAt.After(session.CreatedAt) {
				t.Errorf("ExpiresAt should be after CreatedAt")
			}

			// Should expire in approximately 24 hours
			expectedExpiry := session.CreatedAt.Add(SessionDuration)
			diff := session.ExpiresAt.Sub(expectedExpiry)
			if diff > time.Second || diff < -time.Second {
				t.Errorf("ExpiresAt should be ~24 hours from CreatedAt")
			}
		})
	}
}

func TestSession_IsExpired(t *testing.T) {
	session, _ := NewSession("john", RoleUser)

	// Fresh session should not be expired
	if session.IsExpired() {
		t.Errorf("fresh session should not be expired")
	}

	// Manually expire the session
	session.ExpiresAt = time.Now().Add(-1 * time.Hour)

	if !session.IsExpired() {
		t.Errorf("session with past expiry should be expired")
	}
}

func TestSession_IsValid(t *testing.T) {
	session, _ := NewSession("john", RoleUser)

	// Fresh session should be valid
	if !session.IsValid() {
		t.Errorf("fresh session should be valid")
	}

	// Manually expire the session
	session.ExpiresAt = time.Now().Add(-1 * time.Hour)

	if session.IsValid() {
		t.Errorf("expired session should not be valid")
	}
}

func TestSession_Renew(t *testing.T) {
	session, _ := NewSession("john", RoleUser)

	oldExpiresAt := session.ExpiresAt
	time.Sleep(10 * time.Millisecond)

	session.Renew()

	if !session.ExpiresAt.After(oldExpiresAt) {
		t.Errorf("renewed session should have later expiry time")
	}

	// Should still be valid after renewal
	if !session.IsValid() {
		t.Errorf("renewed session should be valid")
	}

	// New expiry should be approximately 24 hours from now
	expectedExpiry := time.Now().Add(SessionDuration)
	diff := session.ExpiresAt.Sub(expectedExpiry)
	if diff > time.Second || diff < -time.Second {
		t.Errorf("renewed ExpiresAt should be ~24 hours from now")
	}
}

func TestGenerateSessionID_Uniqueness(t *testing.T) {
	// Generate multiple session IDs and verify they're unique
	ids := make(map[string]bool)
	iterations := 100

	for i := 0; i < iterations; i++ {
		id, err := generateSessionID()
		if err != nil {
			t.Fatalf("failed to generate session ID: %v", err)
		}

		if ids[id] {
			t.Errorf("generated duplicate session ID: %s", id)
		}

		ids[id] = true
	}

	if len(ids) != iterations {
		t.Errorf("expected %d unique IDs, got %d", iterations, len(ids))
	}
}

func TestGenerateSessionID_Format(t *testing.T) {
	id, err := generateSessionID()
	if err != nil {
		t.Fatalf("failed to generate session ID: %v", err)
	}

	// Should be 64 characters (32 bytes as hex)
	if len(id) != 64 {
		t.Errorf("expected ID length 64, got %d", len(id))
	}

	// Should only contain hex characters
	for _, char := range id {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f')) {
			t.Errorf("ID contains non-hex character: %c", char)
		}
	}
}

func TestSession_ExpiredSessionRenewal(t *testing.T) {
	session, _ := NewSession("john", RoleUser)

	// Expire the session
	session.ExpiresAt = time.Now().Add(-1 * time.Hour)

	if !session.IsExpired() {
		t.Errorf("session should be expired")
	}

	// Renew it
	session.Renew()

	// Should now be valid
	if !session.IsValid() {
		t.Errorf("renewed session should be valid")
	}

	if session.IsExpired() {
		t.Errorf("renewed session should not be expired")
	}
}
