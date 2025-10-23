package middlewares

import (
    "bytes"
    "encoding/json"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
)

type validateRequest struct {
    Token string `json:"token"`
}

type validateResponse struct {
    Sub   string `json:"sub"`
    Role  string `json:"role"`
    Email string `json:"email"`
    // ... add other fields returned by auth-service
}

// JWTMiddleware returns a Fiber middleware that validates JWT via the auth service.
func JWTMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        var token string

        if authHeader != "" {
            parts := strings.Split(authHeader, " ")
            if len(parts) == 2 && parts[0] == "Bearer" {
                token = parts[1]
            }
        }

        if token == "" {
            token = c.Cookies("access_token")
        }

        if token == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
        }

        authURL := os.Getenv("AUTH_VALIDATE_URL") // e.g. "http://auth-service:8080/api/auth/validate"
        if authURL == "" {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "auth service URL not configured"})
        }

        reqBody, _ := json.Marshal(validateRequest{Token: token})
        client := &http.Client{Timeout: 5 * time.Second}
        req, err := http.NewRequest(http.MethodPost, authURL, bytes.NewReader(reqBody))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to build request"})
        }
        req.Header.Set("Content-Type", "application/json")

        resp, err := client.Do(req)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token validation failed"})
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
        }

        var vResp validateResponse
        if err := json.NewDecoder(resp.Body).Decode(&vResp); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "invalid response from auth service"})
        }

        c.Locals("userID", vResp.Sub)
        c.Locals("userRole", vResp.Role)
        c.Locals("userEmail", vResp.Email)

        return c.Next()
    }
}