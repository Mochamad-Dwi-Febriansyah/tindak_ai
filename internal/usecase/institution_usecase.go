package usecase

import (
	"errors"
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
) 

type InstitutionUsecase struct {
	repo domain.InstitutionRepository
}

func NewInstitutionUsecase(repo domain.InstitutionRepository) *InstitutionUsecase{
	return &InstitutionUsecase{repo}
}

func (u *InstitutionUsecase) GetAllInsitutions() ([]domain.Institution, error) {
	return u.repo.GetAll()
}

func (u *InstitutionUsecase) GetByIDInsitution(id uuid.UUID) (*domain.Institution, error) {
	return u.repo.GetByID(id)
}

func (u *InstitutionUsecase) CreateInsitution(institution *domain.Institution) error {
	existing, err := u.repo.GetByEmail(institution.ContactEmail)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existing != nil {
		return domain.ErrInstitutionEmailExists
	}
	return u.repo.Create(institution)
}

func (u *InstitutionUsecase) UpdateInstitution(institution *domain.Institution) error {
	existing, err := u.repo.GetByID(institution.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrInstitutionNotFound
	}
	if institution.ContactEmail != existing.ContactEmail {
		conflict, _ := u.repo.GetByEmail(institution.ContactEmail)
		if conflict != nil {
			return domain.ErrInstitutionEmailExists
		}
	}

	return u.repo.Update(institution)
}

func (u *InstitutionUsecase) DeleteInsitutions(id uuid.UUID) error {
	return u.repo.Delete(id)
}

func (u *InstitutionUsecase) GetInsitutionByEmail(email string) (*domain.Institution, error) {
	return u.repo.GetByEmail(email)
}


func (u *InstitutionUsecase) GetByIDInsitutionRating(id uuid.UUID) (*domain.InstitutionRating, error) {
	return u.repo.GetRatingByID(id)
}

func (u *InstitutionUsecase) AddRating(institutionRating *domain.InstitutionRating) error {
	return u.repo.AddRating(institutionRating)
}
func (u *InstitutionUsecase) UpdateRating(institutionRating *domain.InstitutionRating) error {
	_, err := u.repo.GetRatingByID(institutionRating.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrInstitutionRatingNotFound
	}
	return u.repo.UpdateRating(institutionRating)
}

func (u *InstitutionUsecase) DeleteRating(id uuid.UUID) error {
	return u.repo.DeleteRating(id)
}