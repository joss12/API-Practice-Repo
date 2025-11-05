package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/shopping-list-backend/internal/auth"
	"github.com/shopping-list-backend/internal/handlers"
	"github.com/shopping-list-backend/internal/store"
)

func main() {
	itemStore := store.NewItemStore()
	userStore := store.NewUserStore()
	sessionStore := store.NewSessionStore()

	// Parse templates with custom functions
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
		"add": func(a, b int) int {
			return a + b
		},
	}

	tmpl := template.Must(template.New("").Funcs(funcMap).ParseGlob("web/templates/*.html"))

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userStore, sessionStore, tmpl)
	itemHandler := handlers.NewItemHandler(itemStore, tmpl)
	adminHandler := handlers.NewAdminHandler(userStore, itemStore, sessionStore, tmpl)

	// Initialize middleware
	authMiddleware := auth.NewMiddleware(sessionStore, userStore)

	// Setup router
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/login", authHandler.ShowLogin).Methods("GET")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/logout", authHandler.Logout).Methods("GET")

	// Protected routes (require authentication)
	r.HandleFunc("/", authMiddleware.RequireAuth(itemHandler.List)).Methods("GET")
	r.HandleFunc("/items", authMiddleware.RequireAuth(itemHandler.Create)).Methods("POST")
	r.HandleFunc("/items/{id}/toggle", authMiddleware.RequireAuth(itemHandler.TogglePurchased)).Methods("POST")
	r.HandleFunc("/items/{id}/delete", authMiddleware.RequireAuth(itemHandler.Delete)).Methods("POST")
	r.HandleFunc("/items/{id}/update", authMiddleware.RequireAuth(itemHandler.Update)).Methods("POST")

	// Admin routes (require admin role)
	r.HandleFunc("/admin", authMiddleware.RequireAdmin(adminHandler.Dashboard)).Methods("GET")
	r.HandleFunc("/admin/clear-purchased", authMiddleware.RequireAdmin(adminHandler.ClearPurchased)).Methods("POST")

	// Start server
	port := "8081"
	fmt.Printf("🚀 Shopping List Server starting on http://localhost:%s\n", port)
	fmt.Println("📝 Login credentials:")
	fmt.Println("   Admin: admin / admin123")
	fmt.Println("   User:  john  / john123")
	fmt.Println("   User:  jane  / jane123")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
