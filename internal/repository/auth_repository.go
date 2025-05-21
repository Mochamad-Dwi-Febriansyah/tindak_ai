package repository

import (
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) Register(input *domain.AuthRegisterInput) error {
	user := &domain.Users{
		ID: 	 uuid.New(),
		FullName: input.FullName,
		Email:    input.Email,
		Gender: input.Gender,
		NumberPhone: input.NumberPhone,
		Address: input.Address,
	}
	return r.db.Create(user).Error
} 