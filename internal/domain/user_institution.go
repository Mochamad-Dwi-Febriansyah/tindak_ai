package domain

import ( 
	"github.com/google/uuid"
)



type UserInstitution struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:char(36);not null;index" json:"user_id"`
	InstitutionID uuid.UUID `gorm:"type:char(36);not null;index" json:"institution_id"`
	Position string `gorm:"type:varchar(255)" json:"position"`

	User Users `gorm:"foreignKey:UserID" json:"-"`
	Institution Institution `gorm:"foreignKey:InstitutionID" json:"-"`
}

type UserInstitutionRepository interface {
	GetAll() ([]UserInstitution, error)
	GetByID(id uuid.UUID) (*UserInstitution, error)
	Create(userInstitution *UserInstitution) error
	Update(userInstitution *UserInstitution) error
	Delete(id uuid.UUID) error
}
