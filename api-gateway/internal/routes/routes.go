// -----------------------------------------------------------------------------
// File: routes.go
//
// This file defines the REST route registration for authentication endpoints in
// the API Gateway. It uses the Fiber web framework to group and register routes
// for user registration, login, token validation, and user info retrieval. It also
// applies JWT middleware to protected routes.
//
// Syntax:
// - Uses Go's function syntax to define route registration.
// - Uses Fiber's routing and grouping features.
// - Applies middleware to specific routes.
//
// Purpose:
// - Centralizes route definitions for maintainability.
// - Ensures consistent application of authentication middleware.
// - Maps REST endpoints to handler functions.
// -----------------------------------------------------------------------------

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
}
