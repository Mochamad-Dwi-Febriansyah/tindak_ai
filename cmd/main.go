package main

import (
	"tindak_ai/internal/delivery/http"
	"tindak_ai/config"
	"tindak_ai/internal/repository"
	"tindak_ai/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
  config.InitDB()

  db := config.DB
  userRepo := repository.NewUserRepository(db)
  userUsercase := usecase.NewUserUsecase(userRepo)
  
  r := gin.Default()
  api := r.Group("/api")
  http.NewUserHandler(api, userUsercase)

 
  r.Run() // listen and serve on 0.0.0.0:8080
}