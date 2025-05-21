package http

import (
	"log"
	"os"
	"tindak_ai/internal/domain"
	"tindak_ai/internal/middleware"
	"tindak_ai/internal/request"
	"tindak_ai/internal/usecase"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(router *gin.RouterGroup, uc *usecase.AuthUsecase, jwtSecret string) {
	handler := &AuthHandler{usecase: uc}

	publicAuth := router.Group("/auth")
	{
		publicAuth.POST("/register", handler.Register)
		publicAuth.POST("/login", handler.Login)
	} 

	protectedAuth := router.Group("/auth")
	protectedAuth.Use(middleware.JWTMiddleware(jwtSecret))
	{
		protectedAuth.POST("/logout", handler.Logout)
		protectedAuth.GET("/me", handler.Me)
		protectedAuth.GET("/permissions", handler.GetUserPermissions)
		protectedAuth.GET("/permissions/:action/:resource", handler.HasPermission)
	} 
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req request.AuthRegisterInput

	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}
	input := domain.AuthRegisterInput{
        FullName:    req.FullName,
        Email:       req.Email,
        Password:    req.Password,
        Gender:      req.Gender,
        NumberPhone: req.NumberPhone, 
    }

	if err := h.usecase.Register(&input); err != nil {
		if err == usecase.ErrEmailExists {
			helper.ValidationFieldErrorResponse(c,"Email", err.Error())
			return
		}else{	
			helper.InternalServerErrorResponse(c, err.Error())
			return
		}
	}
	helper.SuccessResponse(c, "user registered successfully", nil)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.AuthLoginInput
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	input := domain.AuthLoginInput{
		Email: req.Email,
		Password: req.Password,
	}

	token, user, err := h.usecase.Login(&input)
	if err != nil {
		if err == usecase.ErrInvalidCredentials {
			helper.ValidationFieldErrorResponse(c, "Email", err.Error())
			return
		} else if err == usecase.ErrVerification {
			helper.ValidationFieldErrorResponse(c, "Email", err.Error())
			return
		} else if err == usecase.ErrInactive {
			helper.ValidationFieldErrorResponse(c, "Email", err.Error())
			return
		} else{
			helper.InternalServerErrorResponse(c, err.Error())
			return
		}
	}
	// response := gin.H{
	// 	"token": token,
	// 	"user": user,
	// }
	c.SetCookie(
		"jwt_token",      // cookie name
		token,            // value
		3600*1,          // maxAge 24 jam (detik)
		"/",              // path
		"",               // domain (kosong = current domain)
		true,             // secure (true kalau pakai https)
		true,             // httpOnly (tidak bisa diakses JS)
	)

	helper.SuccessResponse(c, "User logged in successfully", user)
}

func (h *AuthHandler) Logout(c *gin.Context){
	c.SetCookie(
		"jwt_token",      // cookie name
		"",               // value
		-1,              // maxAge -1 untuk menghapus cookie
		"/",              // path
		"",               // domain (kosong = current domain)
		true,             // secure (true kalau pakai https)
		true,             // httpOnly (tidak bisa diakses JS)
	)
	
	helper.SuccessResponse(c, "User logged out successfully", nil)
}

func (h *AuthHandler) Me(c *gin.Context) {
	userIDRaw, exist := c.Get(domain.ContextKeyUserID)
	log.Print(userIDRaw)
	if !exist {
		helper.UnauthorizedResponse(c, "User ID not found in context")
		return
	}

	userID, ok := userIDRaw.(string)
	if !ok {
		helper.UnauthorizedResponse(c, "Invalid User ID type")
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		helper.BadRequestResponse(c, "Invalid User ID format")
		return
	}

	user, err := h.usecase.GetProfile(uid)
	if err != nil {
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}

	helper.SuccessResponse(c, "User profile retrieved successfully", user)
}

func (h *AuthHandler) HasPermission(c *gin.Context) {
	userIDRaw, exist := c.Get(domain.ContextKeyUserID)
	if !exist {
		helper.UnauthorizedResponse(c, "User ID not found in context")
		return
	}

	userID, ok := userIDRaw.(string)
	if !ok {
		helper.UnauthorizedResponse(c, "Invalid User ID type")
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		helper.BadRequestResponse(c, "Invalid User ID format")
		return
	}

	action := c.Param("action")
	resource := c.Param("resource")

	hasPermission, err := h.usecase.HasPermission(uid, action, resource)
	if err != nil {
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}

	if hasPermission { 
		c.JSON(200, gin.H{"message": true})
	} else { 
		c.JSON(403, gin.H{"message": false})
	}
}

func (h *AuthHandler) GetUserPermissions(c *gin.Context) {
	userIDRaw, exist := c.Get(domain.ContextKeyUserID)
	if !exist {
		helper.UnauthorizedResponse(c, "User ID not found in context")
		return
	}

	userID, ok := userIDRaw.(string)
	if !ok {
		helper.UnauthorizedResponse(c, "Invalid User ID type")
		return
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		helper.BadRequestResponse(c, "Invalid User ID format")
		return
	}

	permissions, err := h.usecase.GetUserPermissions(uid)
	if err != nil {
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}

	appName := os.Getenv("APP_ENV")
	parsedRespose := make([]string, len(permissions))
	for i := range permissions {
		parsedRespose[i] = appName + "."+ permissions[i].Resource + "." + permissions[i].Action
	} 

	c.JSON(200, gin.H{"permissions": parsedRespose})
}