package http

import (
	"fmt"
	"tindak_ai/internal/domain"
	// "tindak_ai/internal/middleware"
	"tindak_ai/internal/request"
	"tindak_ai/internal/usecase"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
	authRepo domain.AuthRepository
}

func NewUserHandler(router *gin.RouterGroup, uc *usecase.UserUsecase,  authRepo domain.AuthRepository) {
	handler := &UserHandler{
		usecase: uc,
		authRepo: authRepo,
	}

	userGroup := router.Group("/users")
	{
		userGroup.GET("", handler.GetAllUsers)
		userGroup.GET("/:id", handler.GetUserByID)
		userGroup.POST("", handler.CreateUser)
		userGroup.PUT("/:id",handler.UpdateUser)
		userGroup.DELETE("/:id", handler.DeleteUser)
		// userGroup.GET("", middleware.Authorize(authRepo, "read", "user"), handler.GetAllUsers)
		// userGroup.GET("/:id", middleware.Authorize(authRepo, "show", "user"),handler.GetUserByID)
		// userGroup.POST("", middleware.Authorize(authRepo, "create", "user"),handler.CreateUser)
		// userGroup.PUT("/:id", middleware.Authorize(authRepo, "update", "user"),handler.UpdateUser)
		// userGroup.DELETE("/:id", middleware.Authorize(authRepo, "delete", "user"),handler.DeleteUser)
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		helper.InternalServerErrorResponse(c, "failed to fetch users") 
		return
	}
	helper.SuccessResponse(c, "users retrieved successfully", users) 
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error()) 
		return 
	}
	user, err := h.usecase.GetByIDUser(id)
	if err != nil {
		helper.NotFoundResponse(c, err.Error()) 
		return 
	} 
	helper.SuccessResponse(c, "users retrieved successfully", user) 
}

func (h *UserHandler) CreateUser(c *gin.Context){
	var req request.UserCreateRequest

	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return 
	}

	existingUser, err := h.usecase.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		helper.ValidationFieldErrorResponse(c, "Email", "email already in use") 
		return
	}  
 
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil { 
		helper.InternalServerErrorResponse(c, "failed to hash passwod")
		return
	}

	hashedPasswordStr := string(hashedPassword) 

	user := domain.Users{
		ID: uuid.New(),
		FullName: req.FullName,
		Email:       req.Email,
        Password:    &hashedPasswordStr,
        Gender:      req.Gender,
        NumberPhone: req.NumberPhone,
        Address:     req.Address,
        Location:    req.Location,
        DeviceInfo:  req.DeviceInfo,
        AuthProvider:req.AuthProvider,
	}

	file, err := c.FormFile("avatar_url")
	if err == nil {
		path := fmt.Sprintf("uploads/user/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			helper.InternalServerErrorResponse(c, "failed to save avatar") 
			return
		}
		user.AvatarUrl = &path
	}

	if err := h.usecase.CreateUser(&user); err != nil { 
		helper.InternalServerErrorResponse(c,  err.Error())
		return
	} 
	helper.CreatedResponse(c, "user created successfully", user)
}

func (h *UserHandler) UpdateUser(c *gin.Context){
	idParams := c.Param("id")
	userId, err := uuid.Parse(idParams)
	if err != nil { 
		helper.BadRequestResponse(c, "invalid user ID")
		return
	}

	existingUser, err := h.usecase.GetByIDUser(userId)
	if err != nil { 
		helper.NotFoundResponse(c, "user not found")
		return
	}
 
	var req request.UserUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}
 
	if req.Email != nil && *req.Email != existingUser.Email {
		userWithEmail, err := h.usecase.GetUserByEmail(*req.Email)
		if err == nil && userWithEmail != nil && userWithEmail.ID != existingUser.ID {
			helper.ValidationFieldErrorResponse(c, "Email", "email already in use")
			return
		}
	} 
 
	if req.FullName != nil {
		existingUser.FullName = *req.FullName
	}
	if req.Email != nil {
		existingUser.Email = *req.Email
	}
	if req.Gender != nil {
		existingUser.Gender = *req.Gender
	}
	if req.NumberPhone != nil {
		existingUser.NumberPhone = *req.NumberPhone
	}
	if req.Address != nil {
		existingUser.Address = *req.Address
	}
	if req.Location != nil {
		existingUser.Location = req.Location
	}
	if req.DeviceInfo != nil {
		existingUser.DeviceInfo = req.DeviceInfo
	}
	if req.AuthProvider != nil {
		existingUser.AuthProvider = req.AuthProvider
	}

	if req.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			helper.InternalServerErrorResponse(c, "failed to hash password")
			return
		}
		hashedStr := string(hashedPassword)
		existingUser.Password = &hashedStr
	}

	file, err := c.FormFile("avatar_url")
	if err == nil {
		path := fmt.Sprintf("uploads/user/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil { 
			helper.InternalServerErrorResponse(c, "failed to save avatar")
			return
		}
		existingUser.AvatarUrl = &path
	}

	if err := h.usecase.UpdateUser(existingUser); err != nil {
		helper.InternalServerErrorResponse(c, err.Error()) 
		return
	}
 
	helper.SuccessResponse(c, "user updated successfully",existingUser)
}

func (h *UserHandler) DeleteUser(c *gin.Context){
		idParams := c.Param("id")
	userId, err := uuid.Parse(idParams)
	if err != nil { 
		helper.BadRequestResponse(c, "invalid user ID")
		return
	}

	existingUser, err := h.usecase.GetByIDUser(userId)
	if err != nil { 
		helper.NotFoundResponse(c, "user not found")
		return
	} 
	
	if err := h.usecase.DeleteUser(existingUser.ID); err != nil { 
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}  
	helper.SuccessResponse(c, "user deleted successfully",nil)
}
