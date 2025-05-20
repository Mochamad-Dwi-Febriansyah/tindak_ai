package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenderType = string

const (
	Male GenderType = "male"
	Female GenderType = "female"
	Other GenderType = "other"
)

type Users struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey"`
	FullName string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255);uniqueIndex"`
	Password *string `gorm:"type:varchar(255)"`
	Gender GenderType `gorm:"type:enum('male','female','other')"`
	NumberPhone string `gorm:"type:varchar(20)"`
	Address string `gorm:"type:text"`
	AvatarUrl *string `gorm:"type:varchar(255)"`
	Location *string `gorm:"type:varchar(255)"`
	DeviceInfo *string `gorm:"type:varchar(255)"`
	AuthProvider *string  `gorm:"type:varchar(50)"`
	IsVerified bool `gorm:"default:false"`           
	IsActive bool `gorm:"default:true"`         
	LastLoginAt *time.Time `gorm:"type:datetime"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserRepository interface {
	GetAll() ([]Users, error)
	Create(user *Users) error
	GetByID(id uuid.UUID) (*Users, error)
	Update(user *Users) error
	Delete(id uuid.UUID) error
	GetByEmail(email string) (*Users, error)
}