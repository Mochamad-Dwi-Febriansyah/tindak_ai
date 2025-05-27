package repository

import (
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InstitutionRepository struct {
	db *gorm.DB
}

func NewInstitutionRepository(db *gorm.DB) *InstitutionRepository{
	return &InstitutionRepository{db}
}

func (r *InstitutionRepository) GetAll() ([]domain.Institution, error) {
	var institution []domain.Institution
	if err := r.db.Find(&institution).Error; err != nil {
		return nil, err
	}
	return institution, nil
}

func (r *InstitutionRepository) GetByID(id uuid.UUID) (*domain.Institution, error) {
	var institution domain.Institution
	if err := r.db.First(&institution, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &institution, nil
}

func (r * InstitutionRepository) Create(institution *domain.Institution) error {
	return r.db.Create(institution).Error
}

func (r *InstitutionRepository) Update(institution *domain.Institution) error {
	return r.db.Save(institution).Error
}

func (r *InstitutionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Institution{}, "id = ?", id).Error
}

func (r *InstitutionRepository) GetByEmail(email string) (*domain.Institution, error) {
	var institution domain.Institution
	if err := r.db.First(&institution, "contact_email = ?", email).Error; err != nil {
		return nil, err
	}
	return &institution, nil
}



func (r *InstitutionRepository) GetRatingByID(id uuid.UUID) (*domain.InstitutionRating, error) {
	var institutionRating domain.InstitutionRating
	if err := r.db.First(&institutionRating, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &institutionRating, nil
}

func (r *InstitutionRepository) AddRating(institutionRating *domain.InstitutionRating) error {
	return r.db.Create(institutionRating).Error
}
func (r *InstitutionRepository) UpdateRating(institutionRating *domain.InstitutionRating) error {
	return r.db.Save(institutionRating).Error
}
func (r *InstitutionRepository) DeleteRating(id uuid.UUID) error {
	return r.db.Delete(&domain.InstitutionRating{}, "id = ?", id).Error
}