package store

import (
	"errors"
	"github.com/shopping-list-backend/internal/models"
	"sync"
)

var ErrUserNotFound = errors.New("user not found")

// UserStore manages users
type UserStore struct {
	mu    sync.RWMutex
	users map[string]*models.User
}

// NewUserStore creates a store with hardcoded users
func NewUserStore() *UserStore {
	admin, _ := models.NewUser("admin", "admin123", "Administrator", models.RoleAdmin)
	user1, _ := models.NewUser("john", "john123", "John Doe", models.RoleUser)
	user2, _ := models.NewUser("jane", "jane123", "Jane Smith", models.RoleUser)

	return &UserStore{
		users: map[string]*models.User{
			"admin": admin,
			"john":  user1,
			"jane":  user2,
		},
	}
}

// Get retrieves a user by username
func (s *UserStore) Get(username string) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[username]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// Authenticate validates username and password
func (s *UserStore) Authenticate(username, password string) (*models.User, error) {
	user, err := s.Get(username)
	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// GetAll returns all users (for admin dashboard)
func (s *UserStore) GetAll() []*models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}
