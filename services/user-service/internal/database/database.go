package database

import (
	"os"

	"go-microservices/services/user-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init opens a GORM connection and auto-migrates user-service models.
// Returns the *gorm.DB or an error.
func Init() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// fallback for local dev
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate user-service models
	if err := db.AutoMigrate(&models.User{}, &models.Client{}); err != nil {
		return nil, err
	}

	return db, nil
}
