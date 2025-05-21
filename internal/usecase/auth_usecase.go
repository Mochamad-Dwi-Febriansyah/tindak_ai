package usecase

import (
	"errors" 
	"tindak_ai/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

var ErrEmailExists = errors.New("email already in use")

type AuthUsecase struct {
	repo domain.AuthRepository
	userRepo  domain.UserRepository
}

func NewAuthUsecase(repo domain.AuthRepository, userRepo domain.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
		userRepo: userRepo,
	}
}

func (a *AuthUsecase) Register(input *domain.AuthRegisterInput) error {
	existingUser, err := a.userRepo.GetByEmail(input.Email)
	if err == nil && existingUser != nil {
		return ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	input.Password = string(hashedPassword)

	return a.repo.Register(input)
} 