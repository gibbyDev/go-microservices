// -----------------------------------------------------------------------------
// File: handlers.go
//
// This file implements REST API handlers for authentication endpoints in the API
// Gateway. It uses the Fiber web framework to define HTTP handlers that parse
// incoming requests, call the AuthService gRPC client, and return JSON responses.
//
// Syntax:
// - Uses Go's struct and method syntax to organize handlers.
// - Uses Fiber's context for request/response handling.
// - Parses request bodies into protobuf message types.
// - Handles errors and returns appropriate HTTP status codes.
//
// Purpose:
// - Bridges RESTful HTTP requests from the frontend to gRPC calls to backend services.
// - Provides endpoints for user registration, login, token validation, and user info.
// - Ensures consistent error handling and response formatting.
// -----------------------------------------------------------------------------
package handlers

import (
	"context"
	"net/http"

	"github.com/gibbyDev/go-microservices/api-gateway/internal/clients"
	pb "github.com/gibbyDev/go-microservices/proto/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthClient *clients.AuthClient
}

func NewAuthHandler(authClient *clients.AuthClient) *AuthHandler {
	return &AuthHandler{AuthClient: authClient}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req pb.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	resp, err := h.AuthClient.Register(context.Background(), &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req pb.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	resp, err := h.AuthClient.Login(context.Background(), &req)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *AuthHandler) ValidateToken(c *fiber.Ctx) error {
	var req pb.ValidateTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	resp, err := h.AuthClient.ValidateToken(context.Background(), &req)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *AuthHandler) GetUserInfo(c *fiber.Ctx) error {
	var req pb.GetUserInfoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	resp, err := h.AuthClient.GetUserInfo(context.Background(), &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
