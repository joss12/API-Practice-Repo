package models

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name        string
		username    string
		password    string
		displayName string
		role        Role
		shouldError bool
		expectedErr error
	}{
		{
			name:        "valid admin user",
			username:    "admin",
			password:    "admin123",
			displayName: "Administrator",
			role:        RoleAdmin,
			shouldError: false,
		},
		{
			name:        "valid regular user",
			username:    "john",
			password:    "password123",
			displayName: "John Doe",
			role:        RoleUser,
			shouldError: false,
		},
		{
			name:        "empty display name uses username",
			username:    "jane",
			password:    "password123",
			displayName: "",
			role:        RoleUser,
			shouldError: false,
		},
		{
			name:        "empty username",
			username:    "",
			password:    "password123",
			displayName: "Test",
			role:        RoleUser,
			shouldError: true,
			expectedErr: ErrInvalidUsername,
		},
		{
			name:        "empty password",
			username:    "john",
			password:    "",
			displayName: "John",
			role:        RoleUser,
			shouldError: true,
			expectedErr: ErrInvalidPassword,
		},
		{
			name:        "invalid role",
			username:    "john",
			password:    "password123",
			displayName: "John",
			role:        Role("superadmin"),
			shouldError: true,
			expectedErr: ErrInvalidRole,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUser(tt.username, tt.password, tt.displayName, tt.role)

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				if err != tt.expectedErr {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if user.Username != tt.username {
				t.Errorf("expected username %s, got %s", tt.username, user.Username)
			}

			if user.Role != tt.role {
				t.Errorf("expected role %s, got %s", tt.role, user.Role)
			}

			// Password should be hashed
			if user.PasswordHash == tt.password {
				t.Errorf("password should be hashed, not stored in plain text")
			}

			if user.PasswordHash == "" {
				t.Errorf("password hash should not be empty")
			}

			// Check display name
			expectedDisplayName := tt.displayName
			if expectedDisplayName == "" {
				expectedDisplayName = tt.username
			}
			if user.DisplayName != expectedDisplayName {
				t.Errorf("expected display name %s, got %s", expectedDisplayName, user.DisplayName)
			}
		})
	}
}

func TestUser_CheckPassword(t *testing.T) {
	user, err := NewUser("john", "correctpassword", "John Doe", RoleUser)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{"correct password", "correctpassword", true},
		{"wrong password", "wrongpassword", false},
		{"empty password", "", false},
		{"similar password", "correctPassword", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := user.CheckPassword(tt.password)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestUser_IsAdmin(t *testing.T) {
	adminUser, _ := NewUser("admin", "password", "Admin", RoleAdmin)
	regularUser, _ := NewUser("user", "password", "User", RoleUser)

	if !adminUser.IsAdmin() {
		t.Errorf("admin user should return true for IsAdmin()")
	}

	if regularUser.IsAdmin() {
		t.Errorf("regular user should return false for IsAdmin()")
	}
}

func TestUser_IsUser(t *testing.T) {
	adminUser, _ := NewUser("admin", "password", "Admin", RoleAdmin)
	regularUser, _ := NewUser("user", "password", "User", RoleUser)

	if adminUser.IsUser() {
		t.Errorf("admin user should return false for IsUser()")
	}

	if !regularUser.IsUser() {
		t.Errorf("regular user should return true for IsUser()")
	}
}

func TestValidateRole(t *testing.T) {
	tests := []struct {
		name     string
		role     string
		expected bool
	}{
		{"valid user role", "user", true},
		{"valid admin role", "admin", true},
		{"invalid role", "superadmin", false},
		{"empty role", "", false},
		{"random string", "xyz", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateRole(tt.role)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPasswordHashing(t *testing.T) {
	// Create two users with same password
	user1, _ := NewUser("user1", "samepassword", "User 1", RoleUser)
	user2, _ := NewUser("user2", "samepassword", "User 2", RoleUser)

	// Hashes should be different (bcrypt uses salt)
	if user1.PasswordHash == user2.PasswordHash {
		t.Errorf("password hashes should be different even with same password")
	}

	// But both should validate with correct password
	if !user1.CheckPassword("samepassword") {
		t.Errorf("user1 should validate with correct password")
	}

	if !user2.CheckPassword("samepassword") {
		t.Errorf("user2 should validate with correct password")
	}
}
