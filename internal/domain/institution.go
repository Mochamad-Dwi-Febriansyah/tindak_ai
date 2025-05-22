package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Institution struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`	
	Address string `gorm:"type:text" json:"address"`
	PostalCode string `gorm:"type:varchar(20)" json:"postal_code"`
	ContactPhone string `gorm:"type:varchar(20)" json:"contact_phone"`
	ContactEmail string `gorm:"type:varchar(255)" json:"contact_email"`
	Website *string `gorm:"type:varchar(255)" json:"website"`
	LogoUrl *string `gorm:"type:varchar(255)" json:"logo_url"`
	Fax *string `gorm:"type:varchar(20)" json:"fax"`
	Latitude float64 `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude float64 `gorm:"type:decimal(11,8)" json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type InstitutionRepository interface {
	GetAll() ([]Institution, error)
	GetByID(id uuid.UUID) (*Institution, error)
	Create(Institution *Institution) error
	Update(Institution *Institution) error
	Delete(id uuid.UUID) error
	GetByEmail(email string) (*Institution, error)
}

var ErrInstitutionEmailExists = errors.New("institution email already exists")
var ErrInstitutionNotFound = errors.New("institution not found")