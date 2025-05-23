package request

import (
	"time"

	// "github.com/google/uuid"
)

type NewsArticleCreateRequest struct {
	// Slug         string     `form:"slug" binding:"required,min=3,max=100"`
	Title        string     `form:"title" binding:"required,min=3,max=255"`
	Content      string     `form:"content" binding:"required"` 
	// AuthorID     uuid.UUID  `form:"author_id" binding:"required"` // biasanya ambil dari context
	IsPublished  bool       `form:"is_published"`
	EstimatedAt  *time.Time `form:"estimated_at,omitempty"` // optional
}

type NewsArticleUpdateRequest struct {
	// Slug         *string    `form:"slug,omitempty" binding:"omitempty,min=3,max=100"`
	Title        *string    `form:"title,omitempty" binding:"omitempty,min=3,max=255"`
	Content      *string    `form:"content,omitempty"` 
	IsPublished  *bool      `form:"is_published,omitempty"`
	EstimatedAt  *time.Time `form:"estimated_at,omitempty"`
}
