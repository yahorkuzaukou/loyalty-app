package handlers

import "github.com/gofiber/fiber/v2"

func healthHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func SetupHealthHandlers(group fiber.Router) {
	group.Get("/", healthHandler)
}
