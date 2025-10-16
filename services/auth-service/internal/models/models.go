package models

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Email             string    `gorm:"uniqueIndex;not null"`
	Password          string    `gorm:"not null"`
	VerificationToken string    `json:"-"`
	ResetToken        string    `json:"-"`
	ResetTokenExpiry  time.Time `json:"-"`
}

// TODO: Add database models for auth service
// Example models:
// - User model for authentication
// - Token model for refresh tokens
// - Role model for user roles
