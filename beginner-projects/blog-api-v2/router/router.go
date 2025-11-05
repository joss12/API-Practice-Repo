package router

import (
	"github.com/blog-api-v2/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/posts", handlers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", handlers.GetPostByID).Methods("GET")
	router.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")

	return router
}
