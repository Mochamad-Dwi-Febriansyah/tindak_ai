package usecase

import (
	"context"
	"errors"
	"fmt" 
	// "log"
	"time"
	"tindak_ai/internal/domain"
	service "tindak_ai/internal/usecase/token" 

	// "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	// "golang.org/x/oauth2"
	ggoogle "google.golang.org/api/oauth2/v2"
)

var ErrEmailExists = errors.New("email already in use")
var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrVerification = errors.New("account is not verified")
var ErrInactive = errors.New("account is inactive")

type AuthUsecase struct {
	repo domain.AuthRepository
	userRepo  domain.UserRepository
	jwtSecret string 
	tokenService service.TokenService
	passwordResetRepo domain.PasswordResetRepo
	mailer   domain.Mailer
}

func NewAuthUsecase(repo domain.AuthRepository, userRepo domain.UserRepository, jwtSecret string, tokenService service.TokenService, passwordResetRepo domain.PasswordResetRepo, mailer  domain.Mailer) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
		userRepo: userRepo,
		jwtSecret: jwtSecret,
		tokenService: tokenService,
		passwordResetRepo: passwordResetRepo,
				mailer:            mailer,
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

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"user_id": user.ID,
	// 	"email":   user.Email,
	// 	"exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)).Unix(),
	// })

	// tokenString, err := token.SignedString([]byte(a.jwtSecret))
	// if err != nil {
	// 	return "", nil, fmt.Errorf("failed to sign token: %w", err)
	// }

	tokenString, err := a.tokenService.Generate(user)

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

func (a *AuthUsecase) LoginWithGoogle(ctx context.Context, userInfo *ggoogle.Userinfo) (string, error) {
    user, err := a.repo.FindByEmail(ctx, userInfo.Email)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // register otomatis jika belum ada
            newUser := &domain.Users{
                ID:           uuid.New(),
                FullName:     userInfo.Name,
                Email:        userInfo.Email,
				Gender: "other",
                AvatarUrl:    &userInfo.Picture,
                AuthProvider: ptr("google"),
                IsVerified:   true,
                IsActive:     true,
            }
            if err := a.repo.CreateByGoogle(ctx, newUser); err != nil {
                return "", err
            }
            user = newUser
        } else {
            return "", err
        }
    }

    now := time.Now()
    user.LastLoginAt = &now
    _ = a.repo.UpdateLastLogin(ctx, user.ID, now)

    // generate JWT
    token, err := a.tokenService.Generate(user)
    return token, err
}

func ptr(s string) *string {
    return &s
}


func (a *AuthUsecase) SendResetPasswordEmail(email string) error {
	// 1. Cek user berdasarkan email
	user, err := a.userRepo.GetByEmail(email)
	if err != nil || user == nil {
		return ErrEmailExists
	} 
	// 2. Generate token reset password
	token := uuid.NewString()
	expiry := time.Now().Add(1 * time.Hour) // token berlaku 1 jam

	// 3. Simpan token ke repo
	resetToken := &domain.PasswordResetToken{
		Token:     token,
		Email:     email,
		ExpiresAt: expiry,
	}
	if err := a.passwordResetRepo.SaveToken(resetToken); err != nil {
		return fmt.Errorf("gagal simpan token reset: %w", err)
	}

	// 4. Siapkan body email
	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", token)
	emailBody := fmt.Sprintf(`
		<p>Halo,</p>
		<p>Kamu meminta reset password. Klik link di bawah ini untuk mengganti password kamu:</p>
		<p><a href="%s">%s</a></p>
		<p>Link ini berlaku selama 1 jam.</p>
	`, resetLink, resetLink)

	// 5. Kirim email
	if err := a.mailer.Send(email, "Reset Password", emailBody); err != nil {
		return fmt.Errorf("gagal kirim email reset password: %w", err)
	}

	return nil
}

func (a *AuthUsecase) ResetPassword(token, newPassword string) error {
	// 1. Verifikasi token (cek apakah token ada dan belum expired)
	resetToken, err := a.passwordResetRepo.GetByToken(token)
	if err != nil {
		return domain.ErrResetTokenNotFound
	}

	if resetToken.ExpiresAt.Before(time.Now()) {
		return domain.ErrResetTokenExpired
	}

	// 2. Ambil user dari token (misal dengan email)
	user, err := a.userRepo.GetByEmail(resetToken.Email)
	if err != nil || user == nil {
		return domain.ErrUserNotFound
	}

	// 3. Hash new password (gunakan bcrypt atau metode hash lain)
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = ptr(string(hashedPassword))

	if err := a.userRepo.Update(user); err != nil {
		return domain.ErrUpdatePasswordFailed
	}

	// 5. Hapus token reset supaya tidak bisa dipakai ulang
	if err := a.passwordResetRepo.DeleteToken(token); err != nil {
		// Kalau gagal hapus token, log error tapi jangan gagal total
		// Karena password sudah berhasil direset
		// Bisa juga dikembalikan error kalau ingin strict
	}

	return nil
}
