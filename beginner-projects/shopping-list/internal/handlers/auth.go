package handlers

import (
	"html/template"
	"net/http"

	"github.com/shopping-list-backend/internal/auth"
	"github.com/shopping-list-backend/internal/models"
	"github.com/shopping-list-backend/internal/store"
)

type AuthHandler struct {
	userStore    *store.UserStore
	sessionStore *store.SessionStore
	templates    *template.Template
}

func NewAuthHandler(userStore *store.UserStore, sessionStore *store.SessionStore, tmpl *template.Template) *AuthHandler {
	return &AuthHandler{
		userStore:    userStore,
		sessionStore: sessionStore,
		templates:    tmpl,
	}
}

// ShowLogin displays login page
func (h *AuthHandler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Error": r.URL.Query().Get("error"),
	}
	h.templates.ExecuteTemplate(w, "login.html", data)
}

// Login handles login form submission
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate
	user, err := h.userStore.Authenticate(username, password)
	if err != nil {
		http.Redirect(w, r, "/login?error=Invalid credentials", http.StatusSeeOther)
		return
	}

	// Create session
	session, err := models.NewSession(user.Username, user.Role)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	h.sessionStore.Create(session)

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400, // 24 hours
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout handles user logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Get session
	cookie, err := r.Cookie("session_id")
	if err == nil {
		h.sessionStore.Delete(cookie.Value)
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// CheckAuth returns current user info (for AJAX/API)
func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	// Get user from context (set by middleware)
	user := auth.GetUserFromContext(r.Context())

	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"username":"` + user.Username + `","role":"` + string(user.Role) + `"}`))
}
