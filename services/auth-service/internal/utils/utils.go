package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"time"

	"go-microservices/services/auth-service/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

var (
	accessTokenSecret  []byte
	refreshTokenSecret []byte
)

func init() {
	access := os.Getenv("JWT_ACCESS_SECRET")
	if access == "" {
		access = os.Getenv("JWT_SECRET")
	}
	refresh := os.Getenv("JWT_REFRESH_SECRET")
	if refresh == "" {
		refresh = os.Getenv("JWT_SECRET")
	}

	accessTokenSecret = []byte(access)
	refreshTokenSecret = []byte(refresh)
}

// GenerateJWT generates an access token and refresh token for the provided user.
// The function expects the provided models.Auth (or models.User) to have ID, Email and Role fields.
func GenerateJWT(user models.Auth) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
		"iat":   time.Now().Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString(accessTokenSecret)
	if err != nil {
		log.Println("Error generating access token:", err)
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(refreshTokenSecret)
	if err != nil {
		log.Println("Error generating refresh token:", err)
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

// ValidateJWT parses and validates the provided token string.
// If isRefreshToken is true, the refresh secret is used; otherwise the access secret is used.
func ValidateJWT(tokenString string, isRefreshToken bool) (jwt.MapClaims, error) {
	secret := accessTokenSecret
	if isRefreshToken {
		secret = refreshTokenSecret
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// GenerateRandomToken returns a 128-bit random hex token (32 chars).
func GenerateRandomToken() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("Failed to generate random token")
	}
	return hex.EncodeToString(bytes)
}
