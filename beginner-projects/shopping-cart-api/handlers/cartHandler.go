package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shopping-cart-api/models"
)

var CartItems []models.Item

// POST /cart/add
func AddToCart(c *fiber.Ctx) error {
	var item models.Item

	//1.Parse request body
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	//2. Check if item with ID already exists
	for i, existingItem := range CartItems {
		if existingItem.ID == item.ID {
			CartItems[i].Quantity += item.Quantity //increase quantity
			return c.JSON(fiber.Map{
				"message": "Item quantity updated",
				"cart":    CartItems,
			})
		}
	}

	//3. If item not found, add it
	CartItems = append(CartItems, item)

	return c.JSON(fiber.Map{
		"message": "Item added to cart",
		"cart":    CartItems,
	})

	//return c.SendString("Item added to cart")
}

// POST /cart/remove
func RemoveFromCart(c *fiber.Ctx) error {
	type RemoveRquest struct {
		ID int `json:"id"`
	}
	var req RemoveRquest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	//filter out item by ID
	updatedCart := []models.Item{}
	found := false

	for _, item := range CartItems {
		if item.ID != req.ID {
			updatedCart = append(updatedCart, item)
		} else {
			found = true
		}
	}

	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found in cart",
		})
	}

	CartItems = updatedCart

	return c.JSON(fiber.Map{
		"message": "Item removed",
		"cart":    CartItems,
	})
	//return c.SendString("Item removed from cart")
}

// GET /cart/view
func ViewCart(c *fiber.Ctx) error {
	return c.JSON(CartItems)
}

// Get /cart/total
func GetTotal(c *fiber.Ctx) error {
	var total float64 = 0

	for _, item := range CartItems {
		total += item.Price * float64(item.Quantity)
	}

	return c.JSON(fiber.Map{
		"total": total,
	})
	//return c.SendString("Cart total is ...")
}

// DELETE /cart/clear
func ClearCart(c *fiber.Ctx) error {
	CartItems = []models.Item{} //reset
	return c.SendString("Cart cleared")
}
