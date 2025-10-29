package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init opens a GORM connection for the post-service and returns the *gorm.DB.
// Currently there are no models to auto-migrate; this function prepares the DB
// connection so migrations can be added later.
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

	return db, nil
}
