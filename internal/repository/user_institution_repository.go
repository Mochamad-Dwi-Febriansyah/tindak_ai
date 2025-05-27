package repository

import (
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInstitutionRepository struct {
	db *gorm.DB
}

func NewUserInstitutionRepository(db *gorm.DB) *UserInstitutionRepository {
	return &UserInstitutionRepository{db}
}

func (r *UserInstitutionRepository) GetAll() ([]domain.UserInstitution, error) {
	var userInstitutions []domain.UserInstitution
	if err := r.db.Preload("User").Preload("Institution").Find(&userInstitutions).Error; err != nil {
		return nil, err
	}
	return userInstitutions, nil
}

func (r *UserInstitutionRepository) GetByID(id uuid.UUID) (*domain.UserInstitution, error) {
	var userInstitution domain.UserInstitution
	if err := r.db.Preload("User").Preload("Institution").First(&userInstitution, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &userInstitution, nil
}

func (r *UserInstitutionRepository) Create(userInstitution *domain.UserInstitution) error {
	return r.db.Create(userInstitution).Error
}

func (r *UserInstitutionRepository) Update(userInstitution *domain.UserInstitution) error {
	return r.db.Save(userInstitution).Error
}

func (r *UserInstitutionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.UserInstitution{}, "id = ?", id).Error
}