package http

import (
	"tindak_ai/internal/domain"
	"tindak_ai/internal/request"
	"tindak_ai/internal/usecase"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(router *gin.RouterGroup, uc *usecase.AuthUsecase) {
	handler := &AuthHandler{usecase: uc}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handler.Register)
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