package main

import (
	"log"
	"net/http"

	"github.com/blog-api-v2/router"
)

func main() {
	appRouter := router.SetupRouter()

	log.Println("Server running at http://localhost:3002")
	log.Fatal(http.ListenAndServe(":3002", appRouter))
}
