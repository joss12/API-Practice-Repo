package models

import (
	"os"
	"testing"

	"github.com/auth-system/test"
)

func TestMain(m *testing.M) {
	test.SetupTestDB()
	os.Exit(m.Run())
}

func TestRole(t *testing.T) {
	if RoleUser != "user" {
		t.Errorf("expected 'user', got %s", RoleUser)
	}
	if RoleAdmin != "admin" {
		t.Errorf("expected 'admin', got %s", RoleAdmin)
	}
}
