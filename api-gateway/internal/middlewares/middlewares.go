// -----------------------------------------------------------------------------
// File: middlewares.go
//
// This file defines middleware functions for the API Gateway, specifically a JWT
// authentication middleware using the Fiber web framework. The JWTMiddleware function
// checks for a valid JWT token in the Authorization header and blocks unauthorized
// requests.
//
// Syntax:
// - Uses Go's function and closure syntax to define middleware.
// - Uses environment variables for secret management.
// - Parses JWT tokens and validates them.
//
// Purpose:
// - Protects routes by ensuring only requests with valid JWT tokens are allowed.
// - Provides a reusable authentication layer for REST endpoints.
// -----------------------------------------------------------------------------
package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func JWTMiddleware() fiber.Handler {
	secret := os.Getenv("JWT_SECRET")
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}
		return c.Next()
	}
}