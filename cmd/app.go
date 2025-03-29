package cmd

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/yahorkuzaukou/loyalty-app/pkg/api/handlers"
)

func NewApp() *fiber.App {
	app := fiber.New()
	app.Use(
		jwtware.New(
			jwtware.Config{
				SigningKey: jwtware.SigningKey{
					Key: []byte("somesecret"),
				},
			},
		),
	)

	apiGroup := app.Group("/api")
	apiGroup.Group("/health", handlers.GetHealthHandlers()...)

	return app
}
