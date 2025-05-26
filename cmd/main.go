package main

import (
	"os"
	"tindak_ai/config"
	"tindak_ai/internal/delivery/http"
	"tindak_ai/internal/middleware"
	"tindak_ai/internal/repository"
	"tindak_ai/internal/seed"
	"tindak_ai/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
  config.InitDB()

  db := config.DB

  seed.SeedRoleAndPermission(db)

  loggerRepo := repository.NewLoggerRepository(db)
  loggerUsecase := usecase.NewLoggerUsecase(loggerRepo)

  userRepo := repository.NewUserRepository(db)
  userUsercase := usecase.NewUserUsecase(userRepo)  
  
  jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
    jwtSecret = "dev-secret" // fallback dev
	}

  authRepo := repository.NewAuthRepository(db)
  authUsercase := usecase.NewAuthUsecase(authRepo, userRepo, jwtSecret)

  institutionRepo := repository.NewInstitutionRepository(db)
  institutionUsecase := usecase.NewInstitutionUsecase(institutionRepo)

  userInstitutionRepo := repository.NewUserInstitutionRepository(db)
  userInstitutionUsecase := usecase.NewUserInstitutionUsecase(userInstitutionRepo, userRepo, institutionRepo)

  newsArticleRepo := repository.NewNewsArticleRepository(db)
  newsArticleUsecase := usecase.NewNewsArticleUsecase(newsArticleRepo)

  complaintRepo := repository.NewComplaintRepository(db)
  complaintUsecase := usecase.NewComplaintUsecase(complaintRepo, userRepo, institutionRepo)
  
  r := gin.Default()
  api := r.Group("/api")
  
  http.NewAuthHandler(api, authUsercase, loggerUsecase, jwtSecret)

  api.Use(middleware.JWTMiddleware(jwtSecret))

  http.NewUserHandler(api, userUsercase, authRepo, loggerUsecase)
  http.NewInstitutionHandler(api, institutionUsecase, authRepo, loggerUsecase)
  http.NewUserInstitutionHandler(api, userInstitutionUsecase, authRepo, loggerUsecase)
  http.NewNewsArticleHandler(api, newsArticleUsecase, authRepo, loggerUsecase)
  http.NewComplaintHandler(api, complaintUsecase, authRepo, loggerUsecase)



 
  r.Run()  
}