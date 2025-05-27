package usecase

import (
	"errors"
	"fmt"

	// "log"
	"time"
	"tindak_ai/internal/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailExists = errors.New("email already in use")
var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrVerification = errors.New("account is not verified")
var ErrInactive = errors.New("account is inactive")

type AuthUsecase struct {
	repo domain.AuthRepository
	userRepo  domain.UserRepository
	jwtSecret string
}

func NewAuthUsecase(repo domain.AuthRepository, userRepo domain.UserRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
		userRepo: userRepo,
		jwtSecret: jwtSecret,
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

func (a *AuthUsecase) Login(input *domain.AuthLoginInput) (string, *domain.Users, error) {
	// log.Println("Login input:", input)
	user, err := a.repo.Login(input)
	if err != nil {
		return "", nil,  ErrInvalidCredentials
	} 

	if !user.IsVerified {
		return "", nil, ErrVerification
	}

	// Cek apakah akun aktif
	if !user.IsActive {
		return "", nil, ErrInactive
	}
	// log.Println("User found:", user)

	if user == nil || user.Password == nil {
		return "", nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(input.Password)); err != nil {
		return "", nil, ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", nil, fmt.Errorf("failed to sign token: %w", err)
	}

	now := time.Now()
	user.LastLoginAt = &now
	if err := a.userRepo.Update(user); err != nil {
		return "", nil, fmt.Errorf("failed to update user last login: %w", err)
	}

	return tokenString, user, nil
}

func (a *AuthUsecase) GetProfile(userId uuid.UUID) (*domain.Users, error) {
	return a.userRepo.GetByID(userId)
}

func (a *AuthUsecase) HasPermission(userId uuid.UUID, action string, resource string) (bool, error) {
	_, err := a.userRepo.GetByID(userId)
	if err != nil {
		return false, fmt.Errorf("failed to get user: %w", err)
	} 
	return a.repo.HasPermission(userId, action, resource)
}

func (a *AuthUsecase) GetUserPermissions(userID uuid.UUID) ([]domain.Permission, error) {
	return a.repo.GetPermissionsByUserID(userID)
}
