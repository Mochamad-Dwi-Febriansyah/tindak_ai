package main

import (
	"os"
	"tindak_ai/config"
	"tindak_ai/internal/delivery/http"
	"tindak_ai/internal/middleware"
	"tindak_ai/internal/repository"
	"tindak_ai/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
  config.InitDB()

  db := config.DB
  userRepo := repository.NewUserRepository(db)
  userUsercase := usecase.NewUserUsecase(userRepo)

  authRepo := repository.NewAuthRepository(db)

  jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret" // fallback dev
	}

  authUsercase := usecase.NewAuthUsecase(authRepo, userRepo, jwtSecret)
  
  r := gin.Default()
  api := r.Group("/api")
  
  http.NewAuthHandler(api, authUsercase, jwtSecret)

  api.Use(middleware.JWTMiddleware(jwtSecret))
  http.NewUserHandler(api, userUsercase)


 
  r.Run()  
}