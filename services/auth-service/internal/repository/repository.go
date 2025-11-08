package repository

import (
	"go-microservices/services/auth-service/internal/models"

	"gorm.io/gorm"
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

func (r *Repository) GetAuthByEmail(email string) (*models.Auth, error) {
	var a models.Auth
	if err := r.DB.Where("email = ?", email).First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *Repository) GetAuthByID(id uint) (*models.Auth, error) {
	var a models.Auth
	if err := r.DB.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *Repository) CreateTest(t *models.Test) error {
	return r.DB.Create(t).Error
}

func (r *Repository) ListTests() ([]models.Test, error) {
	var tests []models.Test
	if err := r.DB.Order("id desc").Find(&tests).Error; err != nil {
		return nil, err
	}
	return tests, nil
}
