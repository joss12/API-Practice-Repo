package handlers

import (
	"github.com/blog-api/models"
	"github.com/blog-api/utils"
	"github.com/gofiber/fiber/v2"
	"slices"
)

var posts = []models.BlogPost{}

// Create a posts
func CreatePost(c *fiber.Ctx) error {
	var post models.BlogPost
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	post.ID = utils.GenerateID()
	posts = append(posts, post)
	return c.Status(201).JSON(post)
}

// Get all posts (optionally filter author or tag)
func GetAllPosts(c *fiber.Ctx) error {
	author := c.Query("author")
	tag := c.Query("tag")

	if author != "" || tag != "" {
		var filtered []models.BlogPost
		for _, p := range posts {
			if author != "" && p.Author == author {
				filtered = append(filtered, p)
			}

			if tag != "" {

				if slices.Contains(p.Tags, tag) {
					filtered = append(filtered, p)
				}
			}
		}
		return c.JSON(filtered)
	}
	return c.JSON(posts)
}

// Get posts by ID
func GetPostByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, p := range posts {
		if p.ID == id {
			return c.JSON(p)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
}

// Update posts by ID
func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var updated models.BlogPost
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	for i, p := range posts {
		if p.ID == id {
			updated.ID = id
			posts[i] = updated
			return c.JSON(updated)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
}

// Delete posts by ID
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			return c.SendStatus(204)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
}

