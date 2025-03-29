package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yahorkuzaukou/loyalty-app/pkg/api/handlers"
)

func NewApp() *fiber.App {
	app := fiber.New()

	apiGroup := app.Group("/api")
	apiGroup.Group("/health", handlers.GetHealthHandlers()...)

	return app
}
