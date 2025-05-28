package service

import "tindak_ai/internal/domain"

type TokenService interface {
	Generate(user *domain.Users) (string, error)
}
