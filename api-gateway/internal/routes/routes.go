package routes

import (
	"go-microservices/api-gateway/internal/handlers"
	"go-microservices/api-gateway/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	api := app.Group("/api/v1")

	// Health route
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Gateway is running!")
	})

	api.Post("/signup", authHandler.SignUp)
	api.Post("/signin", authHandler.SignIn)
	api.Post("/validate", authHandler.ValidateToken)
	api.Post("/userinfo", middlewares.JWTMiddleware(), authHandler.GetUserInfo)
	// Test endpoints
	api.Post("/test", authHandler.CreateTest)
	api.Get("/tests", authHandler.ListTests)
}
