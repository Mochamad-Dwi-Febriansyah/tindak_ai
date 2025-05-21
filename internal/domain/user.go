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
	ID           uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	FullName     string         `gorm:"type:varchar(255)" json:"full_name"`
	Email        string         `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password     *string        `gorm:"type:varchar(255)" json:"-"`
	Gender       GenderType     `gorm:"type:enum('male','female','other')" json:"gender"`
	NumberPhone  string         `gorm:"type:varchar(20)" json:"number_phone"`
	Address      string         `gorm:"type:text" json:"address"`
	AvatarUrl    *string        `gorm:"type:varchar(255)" json:"avatar_url"`
	Location     *string        `gorm:"type:varchar(255)" json:"location"`
	DeviceInfo   *string        `gorm:"type:varchar(255)" json:"device_info"`
	AuthProvider *string        `gorm:"type:varchar(50)" json:"auth_provider"`
	IsVerified   bool           `gorm:"default:false" json:"is_verified"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	LastLoginAt  *time.Time     `gorm:"type:datetime" json:"last_login_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	UserPermissions []UserPermission `gorm:"foreignKey:UserID" json:"-"`
}

type UserRepository interface {
	GetAll() ([]Users, error)
	Create(user *Users) error
	GetByID(id uuid.UUID) (*Users, error)
	Update(user *Users) error
	Delete(id uuid.UUID) error
	GetByEmail(email string) (*Users, error)
}