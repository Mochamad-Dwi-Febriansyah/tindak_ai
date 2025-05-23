package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsArticle struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Slug string `gorm:"type:varchar(100)" json:"slug"`	
	Title     string         `gorm:"type:varchar(255)" json:"title"`
	Content      string         `gorm:"type:text" json:"content"`
	ThumbnailUrl    *string        `gorm:"type:varchar(255)" json:"thumbnail_url"`
	AuthorID uuid.UUID `gorm:"type:char(36);not null;index" json:"author_id"`
	IsPublished bool 
	EstimatedAt *time.Time `gorm:"type:datetime;index" json:"estimated_at,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"` 
}

type NewsArticleRepository interface {
	GetAll(filter NewsArticleFilter) ([]NewsArticle, error)
	GetByID(id uuid.UUID) (*NewsArticle, error)
	Create(newsArticle *NewsArticle) error
	Update(newsArticle *NewsArticle) error
	Delete(id uuid.UUID) error 
}

type NewsArticleFilter struct {
	Slug string
}

var ErrNewsArticleNotFound = errors.New("news article not found")