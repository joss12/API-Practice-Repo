package handlers

import (
	"context"
	"time"

	"github.com/chat-api/internal/config"
	"github.com/chat-api/internal/models"
	"github.com/chat-api/internal/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	users *repo.UserRepo
	cfg   config.Config
}

func NewAuthHandler(u *repo.UserRepo, cfg config.Config) *AuthHandler {
	return &AuthHandler{users: u, cfg: cfg}
}

func hashPassword(pw string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(b), err
}

func checkPassword(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func (a *AuthHandler) createToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(a.cfg.JWTSecret))
}

func (a *AuthHandler) Register(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	if body.Username == "" || body.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing fields")
	}

	exists, err := a.users.Exists(context.Background(), body.Username)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if exists {
		return fiber.NewError(fiber.StatusConflict, "username taken")
	}

	hash, err := hashPassword(body.Password)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	u := &models.User{
		Username:     body.Username,
		PasswordHash: hash,
	}

	if err := a.users.Create(context.Background(), u); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": u.ID})
}

func (a *AuthHandler) Login(c *fiber.Ctx) error { // ⭐ POINTER RECEIVER
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	if body.Username == "" || body.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "username and password required")
	}

	u, err := a.users.FindByUsername(context.Background(), body.Username)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if err := checkPassword(u.PasswordHash, body.Password); err != nil {
		return fiber.ErrUnauthorized
	}

	token, err := a.createToken(u.ID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
