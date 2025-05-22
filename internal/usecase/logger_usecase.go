package usecase

import (
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
)

type LoggerUsecase struct {
	repo domain.LoggerRepository
}

func NewLoggerUsecase(repo domain.LoggerRepository) *LoggerUsecase {
	return &LoggerUsecase{repo}
}

func (u * LoggerUsecase) Log(method domain.MethodType,level domain.LogLevel, message string, source string, userID *uuid.UUID, userIP string, userAgent string, metadata *string) error {
	log := &domain.Logging{
		Method:   method,
		Level:     level,
		Message:   message,
		Source:   source,
		UserID:    userID,
		UserIP:    &userIP,
		UserAgent: &userAgent, 
		Metadata:  metadata,
	}
	return u.repo.CreateLog(log)
}