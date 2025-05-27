package usecase

import (
	"errors"
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComplaintUsecase struct {
	repo domain.ComplaintRepository
	userRepo domain.UserRepository
	institutionRepo domain.InstitutionRepository
}

func NewComplaintUsecase(repo domain.ComplaintRepository, userRepo domain.UserRepository, institutionRepo domain.InstitutionRepository) *ComplaintUsecase {
	return &ComplaintUsecase{
		repo: repo,
		userRepo: userRepo,
		institutionRepo: institutionRepo,
}
}

func (u *ComplaintUsecase) GetAllComplaint() ([]domain.Complaint, error) {
	return u.repo.GetAll()
}

func (u *ComplaintUsecase) GetComplaintByID(id uuid.UUID) (*domain.Complaint, error) {
	return u.repo.GetByID(id)
}

func (u *ComplaintUsecase) CreateComplaint(complaint *domain.Complaint) error {
	_, err := u.userRepo.GetByID(complaint.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return domain.ErrUserNotFound
		}
		return err
	}
		_, err = u.institutionRepo.GetByID(*complaint.InstitutionID) 
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrInstitutionNotFound
		}
		return err
	} 
	return u.repo.Create(complaint)
}

func (u *ComplaintUsecase) UpdateComplaint(complaint *domain.Complaint) error {
	_, err := u.repo.GetByID(complaint.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrComplaintNotFound
	}
	return u.repo.Update(complaint)
}

func (u *ComplaintUsecase) DeleteComplaint(id uuid.UUID) error {
	return u.repo.Delete(id)
} 

func (u *ComplaintUsecase) GetComplaintByComplaintNumber(id string) (*domain.Complaint, error) {
	return u.repo.GetByComplaintNumber(id)
}

func (u *ComplaintUsecase) GetByIDComplaintRating(id uuid.UUID) (*domain.ComplaintRating, error) {
	return u.repo.GetComplaintRatingByID(id)
}

func (u *ComplaintUsecase) AddComplaintRating(complaintRating *domain.ComplaintRating) error {
	_, err := u.userRepo.GetByID(complaintRating.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return domain.ErrUserNotFound
		}
		return err
	}
	_, err = u.repo.GetByID(complaintRating.ComplaintID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrComplaintNotFound
	}
	return u.repo.AddComplaintRating(complaintRating)
}

func (u *ComplaintUsecase) UpdateComplaintRating(complaintRating *domain.ComplaintRating) error {
	_, err := u.repo.GetComplaintRatingByID(complaintRating.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrComplaintRatingNotFound
	}
	return u.repo.UpdateComplaintRating(complaintRating)
}

func (u *ComplaintUsecase) DeleteComplaintRating(id uuid.UUID) error {
	return u.repo.DeleteComplaintRating(id)
}