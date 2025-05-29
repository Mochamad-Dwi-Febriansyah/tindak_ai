package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type PasswordResetToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Email string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type PasswordResetTokenModel struct {
	Token     string    `gorm:"primaryKey"`
	Email     string
	ExpiresAt time.Time
}

type PasswordResetRepo interface {
	SaveToken(token *PasswordResetToken) error
	GetByToken(token string) (*PasswordResetToken, error)
	DeleteToken(token string) error
}

 

var (
	ErrResetTokenNotFound = errors.New("reset token not found or expired")
	ErrResetTokenExpired  = errors.New("reset token expired") 
	ErrUpdatePasswordFailed = errors.New("gagal update password user")
)