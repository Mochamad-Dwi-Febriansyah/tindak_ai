package repository

import (
	"errors" 
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsArticleRepository struct {
	db *gorm.DB
}

func NewNewsArticleRepository(db *gorm.DB) *NewsArticleRepository {
	return &NewsArticleRepository{db}
}

func (r *NewsArticleRepository) GetAll(filter domain.NewsArticleFilter) ([]domain.NewsArticle, error) {
	var newsArticles []domain.NewsArticle 
	db := r.db 
	if filter.Slug != "" {
		db = db.Where("slug = ?", filter.Slug)
	}
	if err := db.Find(&newsArticles).Error; err != nil {
		return nil, err
	}
	return newsArticles, nil
}

func (r *NewsArticleRepository) GetByID(id uuid.UUID) (*domain.NewsArticle, error) {
	var newsArticle *domain.NewsArticle
	if err := r.db.First(&newsArticle, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNewsArticleNotFound
		}
		return nil , err
	}
	return newsArticle, nil
}

func (r *NewsArticleRepository) Create(newsArticle *domain.NewsArticle) error {
	return r.db.Create(newsArticle).Error
}

func (r *NewsArticleRepository) Update(newsArticle *domain.NewsArticle) error {
	return r.db.Save(newsArticle).Error
}

func (r *NewsArticleRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.NewsArticle{}, "id = ?", id).Error
} 