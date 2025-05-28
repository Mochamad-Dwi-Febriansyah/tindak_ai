package service

import (
	"time"
	"tindak_ai/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secret string
}

func NewJWTService(secret string) TokenService {
	return &jwtService{secret: secret}
}

func (s *jwtService) Generate(user *domain.Users) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "tindak-ai",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}
