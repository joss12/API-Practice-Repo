package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gotasker/internal/config"

	"github.com/gotasker/internal/database"
	"github.com/gotasker/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	data := new(models.User)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	data.Password = string(hash)

	if data.Role == "" {
		data.Role = "user"
	}

	database.DB.Create(&data)
	return c.JSON(fiber.Map{"message": "user created"})
}

func Login(c *fiber.Ctx) error {
	var req models.User
	var user models.User

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	database.DB.Where("email = ?", req.Email).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))

	return c.JSON(fiber.Map{"token": t})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Logged out"})
}
