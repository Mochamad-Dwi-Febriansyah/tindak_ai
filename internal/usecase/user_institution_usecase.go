package usecase

import (
	"errors"
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)



type UserInstitutionUsecase struct {
	repo domain.UserInstitutionRepository
	userRepo domain.UserRepository
	institutionRepo domain.InstitutionRepository
}

func NewUserInstitutionUsecase(userInstitutionRepository domain.UserInstitutionRepository, userRepo domain.UserRepository, institutionRepo domain.InstitutionRepository) *UserInstitutionUsecase {
	return &UserInstitutionUsecase{
		repo: userInstitutionRepository,
		userRepo: userRepo,
		institutionRepo: institutionRepo,
	}
}

func (u *UserInstitutionUsecase) GetAll() ([]domain.UserInstitution, error) {
	userInstitutions, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return userInstitutions, nil
}

func (u *UserInstitutionUsecase) GetByID(id uuid.UUID) (*domain.UserInstitution, error) {
	userInstitution, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return userInstitution, nil
}

func (u *UserInstitutionUsecase) Create(userInstitution *domain.UserInstitution) error {
	_, err := u.userRepo.GetByID(userInstitution.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return domain.ErrUserNotFound
		}
		return err
	}

	_, err = u.institutionRepo.GetByID(userInstitution.InstitutionID) 
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrInstitutionNotFound
		}
		return err
	} 

	err = u.repo.Create(userInstitution)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserInstitutionUsecase) Update(userInstitution *domain.UserInstitution) error {
	_, err := u.userRepo.GetByID(userInstitution.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return domain.ErrUserNotFound
		}
		return err
	}

	_, err = u.institutionRepo.GetByID(userInstitution.InstitutionID) 
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.ErrInstitutionNotFound
		}
		return err
	} 

	err = u.repo.Update(userInstitution)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserInstitutionUsecase) Delete(id uuid.UUID) error {
	err := u.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}