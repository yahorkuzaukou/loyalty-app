package handlers

import "github.com/gofiber/fiber/v2"

func HealthHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetHealthHandlers() []func(c *fiber.Ctx) error {
	return []func(c *fiber.Ctx) error{
		HealthHandler,
	}
}
