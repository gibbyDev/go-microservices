package models

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Username          string    `gorm:"uniqueIndex;not null"`
	Email             string    `gorm:"uniqueIndex;not null"`
	Password          string    `gorm:"not null"`
	VerificationToken string    `json:"-"`
	ResetToken        string    `json:"-"`
	ResetTokenExpiry  time.Time `json:"-"`
	Role              string    `gorm:"type:varchar(50);default:user" json:"role,omitempty"`
}

type Test struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
}
