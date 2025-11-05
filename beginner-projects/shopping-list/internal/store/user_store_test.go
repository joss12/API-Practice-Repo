package store

import (
	"testing"
)

func TestUserStore_Hardcoded(t *testing.T) {
	store := NewUserStore()

	//Should have 3 users
	users := store.GetAll()
	if len(users) != 3 {
		t.Errorf("Expected 3 users, go %d", len(users))
	}

	//Get admin
	admin, err := store.Get("admin")
	if err != nil || !admin.IsAdmin() {
		t.Error("Admin user should exist and be admin")
	}

	//Get regular user
	john, err := store.Get("john")
	if err != nil || !john.IsUser() {
		t.Error("John should exist and be regular user")
	}
}

func TestUserStore_Authenticate(t *testing.T) {
	store := NewUserStore()

	//Valid login
	user, err := store.Authenticate("admin", "admin123")
	if err != nil || user == nil {
		t.Error("Valid credentials should authenticate")
	}

	//Invalid password
	_, err = store.Authenticate("admin", "wrongpas")
	if err == nil {
		t.Error("Invalid password should fail")
	}

	//Invalid user
	_, err = store.Authenticate("nobody", "pass")
	if err != ErrUserNotFound {
		t.Error("Non-existent user should return not found")
	}
}
