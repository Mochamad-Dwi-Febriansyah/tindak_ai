package main

import (
	"os"
	"tindak_ai/config"
	"tindak_ai/internal/delivery/http"
	"tindak_ai/internal/middleware"
	"tindak_ai/internal/repository"
	"tindak_ai/internal/seed"
	"tindak_ai/internal/usecase"
	service "tindak_ai/internal/usecase/token" 
	"tindak_ai/pkg/mailer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
  config.InitDB()

  config.InitGoogleOAuth()

  serviceJwt := service.NewJWTService(os.Getenv("JWT_SECRET"))

  smtpMailer := &mailer.SMTPMailer{
		Host:     os.Getenv("SMTP_HOST"),     // contoh: smtp.gmail.com
		Port:     os.Getenv("SMTP_PORT"),     // contoh: 587
		Username: os.Getenv("SMTP_USERNAME"), // akun email
		Password: os.Getenv("SMTP_PASSWORD"), // app password
		From:     os.Getenv("SMTP_FROM"),     // email pengirim
	} 

  db := config.DB

  // seed.SeedRoleAndPermission(db)
  seed.SeedInstitution(db)

  loggerRepo := repository.NewLoggerRepository(db)
  loggerUsecase := usecase.NewLoggerUsecase(loggerRepo)

  userRepo := repository.NewUserRepository(db)
  userUsercase := usecase.NewUserUsecase(userRepo)  
  
  jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
    jwtSecret = "dev-secret" // fallback dev
	}

  passwordResetRepo := repository.NewPasswordResetRepository(db)

  authRepo := repository.NewAuthRepository(db)
  authUsercase := usecase.NewAuthUsecase(authRepo, userRepo, jwtSecret, serviceJwt, passwordResetRepo, smtpMailer)

  institutionRepo := repository.NewInstitutionRepository(db)
  institutionUsecase := usecase.NewInstitutionUsecase(institutionRepo)

  userInstitutionRepo := repository.NewUserInstitutionRepository(db)
  userInstitutionUsecase := usecase.NewUserInstitutionUsecase(userInstitutionRepo, userRepo, institutionRepo)

  newsArticleRepo := repository.NewNewsArticleRepository(db)
  newsArticleUsecase := usecase.NewNewsArticleUsecase(newsArticleRepo)

  complaintRepo := repository.NewComplaintRepository(db)
  complaintUsecase := usecase.NewComplaintUsecase(complaintRepo, userRepo, institutionRepo)
  
  r := gin.Default()

  r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"}, // sesuaikan dengan frontend kamu
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
  }))

  api := r.Group("/api")
  
  http.NewAuthHandler(api, authUsercase, loggerUsecase, jwtSecret)

  
  api.Use(middleware.JWTMiddleware(jwtSecret))

  http.NewUserHandler(api, userUsercase, authRepo, loggerUsecase)
  http.NewInstitutionHandler(api, institutionUsecase, authRepo, loggerUsecase)
  http.NewUserInstitutionHandler(api, userInstitutionUsecase, authRepo, loggerUsecase)
  http.NewNewsArticleHandler(api, newsArticleUsecase, authRepo, loggerUsecase)
  http.NewComplaintHandler(api, complaintUsecase, authRepo, loggerUsecase)



  r.Static("/uploads", "./uploads")
 
  r.Run()  
}