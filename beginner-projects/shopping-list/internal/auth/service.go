package auth

import (
	"github.com/shopping-list-backend/internal/models"
	"github.com/shopping-list-backend/internal/store"
)

// Service handles authentication business logic
type Service struct {
	userStore    *store.UserStore
	sessionStore *store.SessionStore
}

// NewService creates an auth service
func NewService(userStore *store.UserStore, sessionStore *store.SessionStore) *Service {
	return &Service{
		userStore:    userStore,
		sessionStore: sessionStore,
	}
}

// Login authenticates user and creates session
func (s *Service) Login(username, password string) (*models.Session, error) {

	//Authentite error
	user, err := s.userStore.Authenticate(username, password)
	if err != nil {
		return nil, err
	}

	//Create session
	session, err := models.NewSession(user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	//Store Session
	if err := s.sessionStore.Create(session); err != nil {
		return nil, err
	}
	return session, nil
}

// Logout destroys a session
func (s *Service) Logout(sessionID string) error {
	return s.sessionStore.Delete(sessionID)
}

// ValidateSession checks if session is valid
func (s *Service) ValidateSession(sessionID string) (*models.Session, error) {
	return s.sessionStore.Get(sessionID)
}

// GetUser retrieves user details
func (s *Service) GetUser(username string) (*models.User, error) {
	return s.userStore.Get(username)
}
