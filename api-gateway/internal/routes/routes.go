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
	"github.com/gofiber/fiber/v2"
	"go-microservices/api-gateway/internal/handlers"
	"go-microservices/api-gateway/internal/middlewares"
)

func RegisterAuthRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	api := app.Group("/api/auth")
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)
	api.Post("/validate", authHandler.ValidateToken)
	api.Post("/userinfo", middlewares.JWTMiddleware(), authHandler.GetUserInfo)
}