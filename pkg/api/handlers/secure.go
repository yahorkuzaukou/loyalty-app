package handlers

import "github.com/gofiber/fiber/v2"

func secureHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"secret": "absolute",
	})
}

func SetupSecureHandlers(group fiber.Router) {
	group.Get("/secret", secureHandler)
}
