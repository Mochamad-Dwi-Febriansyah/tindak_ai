package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusLevel string

const (
	StatusLevelInProgress    StatusLevel = "in_progress"
	StatusLevelPending   StatusLevel = "pending"
	StatusLevelRejected StatusLevel = "rejected"
	StatusLevelDone   StatusLevel = "done"
	StatusLevelWithdrawn  StatusLevel = "withdrawn"
)

type RatingLevel int

const (
	Rating1 RatingLevel = 1
	Rating2 RatingLevel = 2
	Rating3 RatingLevel = 3
	Rating4 RatingLevel = 4
	Rating5 RatingLevel = 5
)

func (r RatingLevel) IsValid() bool {
	return r >= Rating1 && r <= Rating5
}

func (r *RatingLevel) BeforeSave(tx *gorm.DB) error {
	if !r.IsValid() {
		return errors.New("invalid rating value, must be between 1 and 5")
	}
	return nil
} 
func (s StatusLevel) IsValid() bool {
	switch s {
	case StatusLevelInProgress, StatusLevelPending, StatusLevelRejected, StatusLevelDone, StatusLevelWithdrawn:
		return true
	}
	return false
}

func parseEstimatedAt(estimatedAtStr string) *string {
	if estimatedAtStr == "" {
		return nil
	}
	return &estimatedAtStr
}

type Complaint struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ComplaintNumber string `gorm:"type:varchar(50);uniqueIndex;not null" json:"complaint_number"`
	UserID uuid.UUID `gorm:"type:char(36);not null;index" json:"user_id"`
	InstitutionID *uuid.UUID `gorm:"type:char(36);not null;index" json:"institution_id"`

	Title     string         `gorm:"type:varchar(255)" json:"title"`
	Description      string         `gorm:"type:text" json:"description"`
	PhotoUrl    *string        `gorm:"type:varchar(255)" json:"photo_url"`
	Location     *string        `gorm:"type:varchar(255)" json:"location"`
	Latitude  *float64 `gorm:"type:decimal(10,8)" json:"latitude,omitempty"`
	Longitude *float64 `gorm:"type:decimal(11,8)" json:"longitude,omitempty"`

	Status StatusLevel `gorm:"type:varchar(15)" json:"status"`

	VerifiedBy *uuid.UUID `gorm:"type:char(36)" json:"verified_by,omitempty"`

	EstimatedAt *time.Time `gorm:"type:datetime;index" json:"estimated_at,omitempty"` 
	StartedAt         *time.Time     `json:"started_at,omitempty"`                             // waktu mulai dikerjakan
	CompletedAt       *time.Time     `json:"completed_at,omitempty"`  

	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	StatusHistory []ComplaintStatusHistory `gorm:"foreignKey:ComplaintID" json:"status_history,omitempty"`
	ComplaintRatings []ComplaintRating `gorm:"foreignKey:ComplaintID" json:"ratings,omitempty"`

}
 

type ComplaintStatusHistory struct {
	ID          uuid.UUID   `gorm:"type:char(36);primaryKey" json:"id"`
	ComplaintID uuid.UUID   `gorm:"type:char(36);not null;index" json:"complaint_id"`
	OldStatus   StatusLevel `gorm:"type:varchar(15)" json:"old_status"`
	NewStatus   StatusLevel `gorm:"type:varchar(15)" json:"new_status"`
	ChangedBy   *uuid.UUID  `gorm:"type:char(36)" json:"changed_by,omitempty"`  
	Note        *string     `gorm:"type:text" json:"note,omitempty"`          // opsional: alasan perubahan
	ChangedAt   time.Time   `gorm:"autoCreateTime" json:"changed_at"`

	Complaint Complaint `gorm:"foreignKey:ComplaintID" json:"-"`
} 

type ComplaintRating struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ComplaintID uuid.UUID `gorm:"type:char(36);not null;index" json:"complaint_id"`
	UserID      uuid.UUID `gorm:"type:char(36);not null;index" json:"user_id"` // siapa yang memberi rating
	Rating      RatingLevel       `gorm:"type:int;not null" json:"rating"`              // misal 1-5
	Comment     *string   `gorm:"type:text" json:"comment,omitempty"`           // opsional, komentar rating
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	
	Complaint Complaint `gorm:"foreignKey:ComplaintID" json:"-"`
	User      Users     `gorm:"foreignKey:UserID" json:"-"`
}


type ComplaintRepository interface {
	GetAll() ([]Complaint, error)
	GetByID(id uuid.UUID) (*Complaint, error)
	Create(complaint *Complaint) error
	Update(complaint *Complaint) error
	Delete(id uuid.UUID) error

	GetByComplaintNumber(cN string) (*Complaint, error)
 
	GetComplaintRatingByID(idComplaintRating uuid.UUID) (*ComplaintRating, error)
	AddComplaintRating(complaintRating *ComplaintRating) error
	UpdateComplaintRating(complaintRating *ComplaintRating) error
	DeleteComplaintRating(idComplaintRating uuid.UUID) error
}

var ErrComplaintNotFound = errors.New("complaint not found")
var ErrComplaintRatingNotFound = errors.New("complaint rating not found")

