package domain

import ( 

	"github.com/google/uuid"
)

type Role struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`

	RolePermission []RolePermission `gorm:"foreignKey:RoleID" json:"-"` 
}

type Permission struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"` 
	Action string `gorm:"type:varchar(255);not null" json:"action"`
	Resource string `gorm:"type:varchar(255);not null" json:"resource"`

	RolePermission []RolePermission `gorm:"foreignKey:PermissionID" json:"-"`
	UserPermission []UserPermission `gorm:"foreignKey:PermissionID" json:"-"`
}

type RolePermission struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	RoleID       uuid.UUID `gorm:"type:char(36);index"`
	PermissionID uuid.UUID `gorm:"type:char(36);index"`

	Role Role `gorm:"foreignKey:RoleID" json:"-"`
	Permission Permission `gorm:"foreignKey:PermissionID" json:"-"`
}

type UserPermission struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID       uuid.UUID `gorm:"type:char(36);index"`
	PermissionID uuid.UUID `gorm:"type:char(36);index"`

	User Users `gorm:"foreignKey:UserID" json:"-"`
	Permission Permission `gorm:"foreignKey:PermissionID" json:"-"`
}
