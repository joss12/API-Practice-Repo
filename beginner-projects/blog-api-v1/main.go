package main

import (
	"github.com/blog-api-v1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//Register routes
	routes.AppRoutes(server)

	server.Run(":3002")
}
