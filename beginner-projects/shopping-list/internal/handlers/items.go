package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shopping-list-backend/internal/auth"
	"github.com/shopping-list-backend/internal/models"
	"github.com/shopping-list-backend/internal/store"
)

type ItemHandler struct {
	itemStore *store.ItemStore
	templates *template.Template
}

func NewItemHandler(itemStore *store.ItemStore, tmpl *template.Template) *ItemHandler {
	return &ItemHandler{
		itemStore: itemStore,
		templates: tmpl,
	}
}

// List shows all items
func (h *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserFromContext(r.Context())
	items := h.itemStore.GetAll()

	data := map[string]interface{}{
		"User":  user,
		"Items": items,
	}

	h.templates.ExecuteTemplate(w, "items.html", data)
}

// Create adds a new item
func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user := auth.GetUserFromContext(r.Context())

	name := r.FormValue("name")
	category := r.FormValue("category")
	quantityStr := r.FormValue("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		quantity = 1
	}

	item, err := models.NewItem(name, category, user.Username, quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.itemStore.Create(item)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// TogglePurchased marks item as purchased/not purchased
func (h *ItemHandler) TogglePurchased(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	item, err := h.itemStore.Get(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	if item.IsPurchased {
		item.MarkAsNotPurchased()
	} else {
		item.MarkAsPurchased()
	}

	h.itemStore.Update(id, item)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Delete removes an item
func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.itemStore.Delete(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Update modifies an item
func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	item, err := h.itemStore.Get(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	// Update fields
	if name := r.FormValue("name"); name != "" {
		item.Name = name
	}
	if category := r.FormValue("category"); category != "" {
		item.Category = category
	}
	if quantityStr := r.FormValue("quantity"); quantityStr != "" {
		if quantity, err := strconv.Atoi(quantityStr); err == nil {
			item.UpdateQuantity(quantity)
		}
	}

	h.itemStore.Update(id, item)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
