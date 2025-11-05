package handlers

import (
	"html/template"
	"net/http"

	"github.com/shopping-list-backend/internal/auth"
	"github.com/shopping-list-backend/internal/store"
)

type AdminHandler struct {
	userStore    *store.UserStore
	itemStore    *store.ItemStore
	sessionStore *store.SessionStore
	templates    *template.Template
}

func NewAdminHandler(userStore *store.UserStore, itemStore *store.ItemStore, sessionStore *store.SessionStore, tmpl *template.Template) *AdminHandler {
	return &AdminHandler{
		userStore:    userStore,
		itemStore:    itemStore,
		sessionStore: sessionStore,
		templates:    tmpl,
	}
}

// Dashboard shows admin dashboard
func (h *AdminHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserFromContext(r.Context())

	users := h.userStore.GetAll()
	items := h.itemStore.GetAll()

	//Calculate stats
	totalItems := len(items)
	purchasedItems := 0
	for _, item := range items {
		if item.IsPurchased {
			purchasedItems++
		}
	}

	data := map[string]interface{}{
		"User":           user,
		"Users":          users,
		"TotalItems":     totalItems,
		"PurchasedItems": purchasedItems,
		"PendingItems":   totalItems - purchasedItems,
	}

	h.templates.ExecuteTemplate(w, "admin.html", data)
}

// ClearPurchased removes all purchased items (admin only)
func (h *AdminHandler) ClearPurchased(w http.ResponseWriter, r *http.Request) {
	items := h.itemStore.GetAll()

	for _, item := range items {
		if item.IsPurchased {
			h.itemStore.Delete(item.ID)
		}
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
