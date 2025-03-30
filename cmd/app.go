package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yahorkuzaukou/loyalty-app/pkg/api/handlers"

	jwtware "github.com/gofiber/contrib/jwt"
)

func NewApp() *fiber.App {
	app := fiber.New()

	apiGroup := app.Group("/api")
	healthGroup := apiGroup.Group("/health")
	handlers.SetupHealthHandlers(healthGroup)
	authGroup := apiGroup.Group("/auth")
	handlers.SetupAuthHandlers(authGroup)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	secureGroup := apiGroup.Group("/secure")
	handlers.SetupSecureHandlers(secureGroup)

	return app
}
