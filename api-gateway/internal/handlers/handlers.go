package handlers

import (
	"context"
	"net/http"

	"go-microservices/api-gateway/internal/clients"
	pb "go-microservices/proto/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthClient *clients.AuthClient
}

func NewAuthHandler(authClient *clients.AuthClient) *AuthHandler {
	return &AuthHandler{AuthClient: authClient}
}

func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	var req pb.SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	resp, err := h.AuthClient.SignUp(context.Background(), &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	var req pb.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	resp, err := h.AuthClient.SignIn(context.Background(), &req)
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

// CreateTest forwards a test creation request to the auth service
func (h *AuthHandler) CreateTest(c *fiber.Ctx) error {
	var body struct {
		Test string `json:"test"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	req := pb.CreateTestRequest{Content: body.Test}
	resp, err := h.AuthClient.CreateTest(context.Background(), &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

// ListTests retrieves tests from the auth service
func (h *AuthHandler) ListTests(c *fiber.Ctx) error {
	resp, err := h.AuthClient.ListTests(context.Background(), &pb.ListTestsRequest{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}
