package usecase

import ( 
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo}
}

func (u *UserUsecase) GetAllUsers() ([]domain.Users, error) {
	return u.repo.GetAll()
}

func (u *UserUsecase) GetByIDUsers(id uuid.UUID) (*domain.Users, error) {
	return u.repo.GetByID(id)
} 

func (u *UserUsecase) CreateUser(user *domain.Users) error { 
	return u.repo.Create(user)
}

func (u *UserUsecase) UpdateUser(user *domain.Users) error {
	return u.repo.Update(user)
}

func (u *UserUsecase) DeleteUser(id uuid.UUID) error {
	return u.repo.Delete(id)
} 
func (u *UserUsecase) GetUserByEmail(email string) (*domain.Users, error) { 
    return u.repo.GetByEmail(email)
}