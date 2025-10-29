package repository

import (
	"errors"

	"gorm.io/gorm"

	"go-microservices/services/auth-service/internal/models"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateAuth(a *models.Auth) error {
	return r.DB.Create(a).Error
}

func (r *Repository) GetByEmail(email string) (*models.Auth, error) {
	var a models.Auth
	if err := r.DB.Where("email = ?", email).First(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

// TODO: Add database repository implementations
// Example repositories:
// - UserRepository for user data access
// - TokenRepository for token management
// - RoleRepository for role management
