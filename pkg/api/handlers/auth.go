package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(c *fiber.Ctx) error {
	data := new(LoginData)

	if err := c.BodyParser(data); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if data.Username != "john" || data.Password != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"name": "username",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": signedToken})
}

func SetupAuthHandlers(group fiber.Router) {
	group.Post("/login", loginHandler)
}
