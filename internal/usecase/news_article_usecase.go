package usecase

import (
	"errors"
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsArticleUsecase struct {
	repo domain.NewsArticleRepository
}

func NewNewsArticleUsecase(repo domain.NewsArticleRepository) *NewsArticleUsecase {
	return &NewsArticleUsecase{
		repo: repo,
	}
}

func (u *NewsArticleUsecase) GetAllNewsArticle(filter domain.NewsArticleFilter) ([]domain.NewsArticle, error){
	return u.repo.GetAll(filter)
}

func (u *NewsArticleUsecase) GetNewsArticleByID(id uuid.UUID) (*domain.NewsArticle, error) {
	return u.repo.GetByID(id)
}

func (u *NewsArticleUsecase) CreateNewsArticle(newsArticle *domain.NewsArticle) error {
	return u.repo.Create(newsArticle)
}

func (u *NewsArticleUsecase) UpdateNewsArticle(newsArticle *domain.NewsArticle) error {
	_, err := u.repo.GetByID(newsArticle.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrNewsArticleNotFound
	}
	return u.repo.Update(newsArticle)
}

func (u *NewsArticleUsecase) DeleteNewsArticle(id uuid.UUID) error {
	return u.repo.Delete(id)
} 