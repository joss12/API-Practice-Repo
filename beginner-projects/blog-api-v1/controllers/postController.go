package controllers

import (
	"net/http"

	"github.com/blog-api-v1/models"
	"github.com/gin-gonic/gin"
)

var posts = []models.Post{
	{ID: "1", Title: "First Post", Content: "Hello, World"},
	{ID: "2", Title: "Gin is great", Content: "Gin makes routing simple."},
}

// GET /posts
func GetPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

// GET /posts/:id
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	for _, p := range posts {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}

// POST /posts
func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	posts = append(posts, newPost)
	c.JSON(http.StatusCreated, newPost)
}

// PUT /posts/:id
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var updated models.Post
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	for i, p := range posts {
		if p.ID == id {
			posts[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}

// DELETE /posts/:id
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}
