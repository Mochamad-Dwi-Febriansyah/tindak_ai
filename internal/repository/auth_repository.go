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

func (r *AuthRepository) HasPermission(userID uuid.UUID, action string, resource string) (bool, error) {
	var count int64
	err := r.db.Table("user_permissions").
		Joins("JOIN permissions ON user_permissions.permission_id = permissions.id").
		Where("user_permissions.user_id = ? AND permissions.action = ? AND permissions.resource = ?", userID, action, resource).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (r *AuthRepository) GetPermissionsByUserID(userID uuid.UUID) ([]domain.Permission, error) {
	var permissions []domain.Permission
	err := r.db.Table("user_permissions").
		Select("permissions.*").
		Joins("JOIN permissions ON user_permissions.permission_id = permissions.id").
		Where("user_permissions.user_id = ?", userID).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}