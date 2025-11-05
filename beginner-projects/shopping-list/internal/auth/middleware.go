package auth

import (
	"context"
	"net/http"

	"github.com/shopping-list-backend/internal/models"
	"github.com/shopping-list-backend/internal/store"
)

type contextKey string

const (
	SessionKey contextKey = "session"
	UserKey    contextKey = "user"
)

// Middleware wraps handlers to require authentication
type Middleware struct {
	sessionStore *store.SessionStore
	userStore    *store.UserStore
}

// NewMiddleware creates auth middleware
func NewMiddleware(sessionStore *store.SessionStore, userStore *store.UserStore) *Middleware {
	return &Middleware{
		sessionStore: sessionStore,
		userStore:    userStore,
	}
}

// RequireAuth ensures user is logged in
func (m *Middleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get session cookie

		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		//Validate session
		session, err := m.sessionStore.Get(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		//Get user
		user, err := m.userStore.Get(session.Username)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		//Add to context
		ctx := context.WithValue(r.Context(), SessionKey, session)
		ctx = context.WithValue(ctx, UserKey, user)

		next(w, r.WithContext(ctx))
	}
}

// RequireAdmin ensures user is admin
func (m *Middleware) RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return m.RequireAuth(func(w http.ResponseWriter, r *http.Request) {
		user := GetUserFromContext(r.Context())

		if !user.IsAdmin() {
			http.Error(w, "Forbidden: Admin only", http.StatusForbidden)
			return
		}

		next(w, r)
	})
}

// Helper functions to get values from context
func GetUserFromContext(ctx context.Context) *models.User {
	user, _ := ctx.Value(UserKey).(*models.User)
	return user
}

func GetSessionFromContext(ctx context.Context) *models.Session {
	session, _ := ctx.Value(SessionKey).(*models.Session)
	return session
}
