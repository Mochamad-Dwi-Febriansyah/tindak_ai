package repository

import (
	"tindak_ai/internal/domain"

	"gorm.io/gorm"
)

type PasswordResetRepository struct {
	db *gorm.DB
}

func NewPasswordResetRepository(db *gorm.DB) *PasswordResetRepository {
	return &PasswordResetRepository{db: db}
}

func (r *PasswordResetRepository) SaveToken(token *domain.PasswordResetToken) error {
	model := domain.PasswordResetTokenModel{
		Token:     token.Token,
		Email:     token.Email,
		ExpiresAt: token.ExpiresAt,
	}
	return r.db.Create(&model).Error
}

func (r *PasswordResetRepository) GetByToken(token string) (*domain.PasswordResetToken, error) {
	var model domain.PasswordResetTokenModel
	if err := r.db.First(&model, "token = ?", token).Error; err != nil {
		return nil, err
	}
	return &domain.PasswordResetToken{
		Token:     model.Token,
		Email:     model.Email,
		ExpiresAt: model.ExpiresAt,
	}, nil
}

func (r *PasswordResetRepository) DeleteToken(token string) error {
	return r.db.Delete(&domain.PasswordResetTokenModel{}, "token = ?", token).Error
}
