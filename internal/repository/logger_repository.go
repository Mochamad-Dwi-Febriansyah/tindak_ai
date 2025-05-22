package repository

import (
	"tindak_ai/internal/domain"

	"gorm.io/gorm"
)

type LoggerRepository struct {
	db *gorm.DB
}

func NewLoggerRepository(db *gorm.DB) *LoggerRepository {
	return &LoggerRepository{db}
}

func (r *LoggerRepository) CreateLog(log *domain.Logging) error {
	return r.db.Create(log).Error
}