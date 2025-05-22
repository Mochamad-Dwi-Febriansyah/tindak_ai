package repository

import ( 
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAll() ([]domain.Users, error) {
	var users []domain.Users
	if err:= r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id uuid.UUID) (*domain.Users, error) {
	var user domain.Users
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *domain.Users) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *domain.Users) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Users{}, "id = ?", id).Error
}

func (r *UserRepository) GetByEmail(email string) (*domain.Users, error) {
	var user domain.Users
    if err := r.db.Unscoped().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
    } 
    return &user, nil
}
