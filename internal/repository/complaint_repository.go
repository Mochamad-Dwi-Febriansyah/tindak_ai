package repository

import (
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComplaintRepository struct {
	db *gorm.DB
}

func NewComplaintRepository(db *gorm.DB) *ComplaintRepository {
	return &ComplaintRepository{db}
}

func (r *ComplaintRepository) GetAll() ([]domain.Complaint, error) {
	var complaints []domain.Complaint
	if err := r.db.Find(&complaints).Error; err != nil {
		return nil, err
	}
	return complaints, nil
} 

func (r *ComplaintRepository) GetByID(id uuid.UUID) (*domain.Complaint, error) {
	var complaint domain.Complaint
	if err := r.db.
		Preload("StatusHistory").
		Preload("ComplaintRatings").
		First(&complaint, "id = ?", id).Error; err != nil{
		return nil, err
	}
	return &complaint, nil
}

func (r *ComplaintRepository) Create(complaint *domain.Complaint) error {
	return r.db.Create(complaint).Error
}

func (r *ComplaintRepository) Update(complaint *domain.Complaint) error {
	return r.db.Save(complaint).Error
}

func (r *ComplaintRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Complaint{}, "id = ?", id).Error
}

func (r *ComplaintRepository) GetByComplaintNumber(idCN string) (*domain.Complaint, error) {
	var complaint domain.Complaint
	if err := r.db.
		Preload("StatusHistory").
		Preload("ComplaintRatings").
		First(&complaint, "complaint_number = ?", idCN).Error; err != nil{
		return nil, err
	}
	return &complaint, nil
}


func (r *ComplaintRepository) GetComplaintRatingByID(idComplaintRating uuid.UUID) (*domain.ComplaintRating, error) {
	var complaintRating domain.ComplaintRating
	if err := r.db.First(&complaintRating, "id = ?", idComplaintRating).Error; err != nil {
		return nil, err
	}
	return &complaintRating, nil
}

func (r *ComplaintRepository) AddComplaintRating(complaintRating *domain.ComplaintRating) error {
	return r.db.Create(complaintRating).Error
}

func (r *ComplaintRepository) UpdateComplaintRating(complaintRating *domain.ComplaintRating) error {
	return r.db.Save(complaintRating).Error
}

func (r *ComplaintRepository) DeleteComplaintRating(id uuid.UUID) error {
	return r.db.Delete(&domain.ComplaintRating{}, "id = ?", id).Error
}