package domain

import "github.com/google/uuid"

type AuthLoginInput struct {
	Email string
	Password string
}

type AuthRegisterInput struct {
	FullName string
	Email string
	Password string
	Gender string
	NumberPhone string
	Address string
}

type AuthResetPasswordInput struct {
	Email string
	NewPassword string
}

type AuthChangePassword struct {
	UserID uuid.UUID
	OldPassword string
	NewPassword string
}

type AuthRepository interface {
	Register(input *AuthRegisterInput) error
	// Login(input *AuthLoginInput) (*Users, error)
	// VerifyEmail(email string) error
	// ResetPassword(input *AuthResetPasswordInput) error
	// ChangePassword(input *AuthChangePassword) error
	// Logout(userID uuid.UUID) error
}