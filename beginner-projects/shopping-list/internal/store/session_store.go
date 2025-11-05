package store

import (
	"sync"

	"github.com/shopping-list-backend/internal/models"
)

// SessionStore manages user sessions in memory
type SessionStore struct {
	mu       sync.RWMutex
	sessions map[string]*models.Session // sessionID -> Session
}

// NewSessionStore creates a new session store
func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]*models.Session),
	}
}

// Create stores a new session
func (s *SessionStore) Create(session *models.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.ID] = session
	return nil
}

// Get retrieves a session by ID
func (s *SessionStore) Get(sessionID string) (*models.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, exists := s.sessions[sessionID]
	if !exists {
		return nil, models.ErrInvalidSession
	}

	// Check if expired
	if session.IsExpired() {
		return nil, models.ErrSessionExpired
	}

	return session, nil
}

// Delete removes a session (logout)
func (s *SessionStore) Delete(sessionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.sessions, sessionID)
	return nil
}

// CleanupExpired removes all expired sessions (call periodically)
func (s *SessionStore) CleanupExpired() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	count := 0
	for id, session := range s.sessions {
		if session.IsExpired() {
			delete(s.sessions, id)
			count++
		}
	}
	return count
}
