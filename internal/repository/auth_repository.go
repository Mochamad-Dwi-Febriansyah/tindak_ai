package repository

import ( 
	// "log"
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) Register(input *domain.AuthRegisterInput) error {
	user := &domain.Users{
		ID: 	 uuid.New(),
		FullName: input.FullName,
		Email:    input.Email,
		Gender: input.Gender,
		Password: &input.Password,
		NumberPhone: input.NumberPhone,
		Address: input.Address,
	}
	return r.db.Create(user).Error
} 

func (r *AuthRepository) Login(input *domain.AuthLoginInput)(*domain.Users, error){
	var user domain.Users
	if err := r.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, err
	}    

	return &user, nil
}