package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name    string
	Email   string `gorm:"uniqueIndex;not null"`
	Active  bool   `gorm:"default:true"`
	Address string
}

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	Name         string
	Role         string `gorm:"default:User"`
	Active       bool   `gorm:"default:true"`
	Username     string `gorm:"uniqueIndex"`
	Address      string
	PhoneNumber  string
	ProfilePhoto []byte `gorm:"type:bytea"`
	ClientID     *uint
	Client       *Client
}
