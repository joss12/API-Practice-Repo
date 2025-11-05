package routes

import (
	"github.com/blog-api-v1/controllers"
	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) {
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPostByID)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
}
