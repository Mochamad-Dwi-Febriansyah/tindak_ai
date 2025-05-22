package http

import (
	"errors"
	"fmt"
	"tindak_ai/internal/domain"
	// "tindak_ai/internal/middleware"
	"tindak_ai/internal/request"
	"tindak_ai/internal/usecase"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InstitutionHandler struct {
	usecase *usecase.InstitutionUsecase
	authRepo domain.AuthRepository
}

func NewInstitutionHandler(router *gin.RouterGroup , uc *usecase.InstitutionUsecase, authRepo domain.AuthRepository) {
	handler := &InstitutionHandler{
		usecase: uc,
		authRepo: authRepo,
	}
	institutionGroup := router.Group("/institutions")
	{
		institutionGroup.GET("" ,handler.GetAll)
		institutionGroup.GET("/:id", handler.GetByID)
		institutionGroup.POST("", handler.Create)
		institutionGroup.PUT("/:id", handler.Update)
		institutionGroup.DELETE("/:id",  handler.Delete)
		institutionGroup.GET("/email",  handler.GetByEmail)
		// institutionGroup.GET("", middleware.Authorize(authRepo, "read", "institution") ,handler.GetAll)
		// institutionGroup.GET("/:id", middleware.Authorize(authRepo, "show", "institution"), handler.GetByID)
		// institutionGroup.POST("", middleware.Authorize(authRepo, "create", "institution"), handler.Create)
		// institutionGroup.PUT("/:id", middleware.Authorize(authRepo, "update", "institution"), handler.Update)
		// institutionGroup.DELETE("/:id", middleware.Authorize(authRepo, "delete", "institution"), handler.Delete)
		// institutionGroup.GET("/email", middleware.Authorize(authRepo, "read_by_email", "institution"), handler.GetByEmail)
	}
}

func (h *InstitutionHandler) GetAll(c *gin.Context) {
	institution, err := h.usecase.GetAllInsitutions()
	if err != nil {
		helper.InternalServerErrorResponse(c, "failed to fetch institutions")
		return
	}
	helper.SuccessResponse(c, "institutions retrieved successfully", institution)
}

func (h *InstitutionHandler) GetByID(c *gin.Context) {
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	} 
	institution, err := h.usecase.GetByIDInsitution(id)
	if err != nil {
		helper.NotFoundResponse(c, err.Error())
		return
	} 
	helper.SuccessResponse(c, "institution retrieved successfully", institution)
}

func (h *InstitutionHandler) Create(c *gin.Context) {
	var req request.InstitutionCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}  

	institution := domain.Institution{
		ID: uuid.New(),
		Name:        req.Name,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
		Address:     req.Address,
		PostalCode: req.PostalCode,
		Website:     req.Website,  
		Fax:	   req.Fax,
		Latitude:  req.Latitude,
		Longitude : req.Longitude,  
	}

	file, err := c.FormFile("logo_url")
	if err == nil {
		path := fmt.Sprintf("uploads/institution/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			helper.InternalServerErrorResponse(c, "failed to save file")
			return
		}
		institution.LogoUrl = &path
	} 

	err = h.usecase.CreateInsitution(&institution)
	if  err != nil {
		if errors.Is(err, domain.ErrInstitutionEmailExists) {
			helper.ValidationFieldErrorResponse(c, "ContactEmail", "email already exists")
			return
		}
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}
	helper.CreatedResponse(c, "institution created successfully", institution)
}

func (h *InstitutionHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		helper.BadRequestResponse(c, "invalid UUID")
		return
	}

	var req request.InstitutionUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	existingInstitution, err := h.usecase.GetByIDInsitution(id)
	if err != nil {
		helper.NotFoundResponse(c, "institution not found")
		return
	}

	if req.Name != nil {
		existingInstitution.Name = *req.Name
	}
	if req.ContactEmail != nil {
		existingInstitution.ContactEmail = *req.ContactEmail
	}
	if req.ContactPhone != nil {
		existingInstitution.ContactPhone = *req.ContactPhone
	}
	if req.Address != nil {
		existingInstitution.Address = *req.Address
	}
	if req.PostalCode != nil {
		existingInstitution.PostalCode = *req.PostalCode
	}
	if req.Website != nil {
		existingInstitution.Website = req.Website
	}
	if req.Fax != nil {
		existingInstitution.Fax = req.Fax
	}
	if req.Latitude != nil {
		existingInstitution.Latitude = *req.Latitude
	}
	if req.Longitude != nil {
		existingInstitution.Longitude = *req.Longitude
	} 

	// Optional file upload
	file, err := c.FormFile("logo_url")
	if err == nil {
		path := fmt.Sprintf("uploads/institutions/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			helper.InternalServerErrorResponse(c, "failed to save file")
			return
		}
		existingInstitution.LogoUrl = &path // or use &path
	}

	err = h.usecase.UpdateInstitution(existingInstitution)
	if err != nil {
		if errors.Is(err, domain.ErrInstitutionEmailExists) {
			helper.ValidationFieldErrorResponse(c, "ContactEmail", "email already exists")
			return
		}
		if errors.Is(err, domain.ErrInstitutionNotFound) {
			helper.NotFoundResponse(c, "institution not found")
			return
		}
		helper.InternalServerErrorResponse(c, "failed to update institution")
		return
	}

	helper.SuccessResponse(c, "institution updated successfully", existingInstitution)
}

func (h *InstitutionHandler) Delete(c *gin.Context) {
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}
	existingInstitution, err := h.usecase.GetByIDInsitution(id)
	if err != nil {
		helper.NotFoundResponse(c, err.Error())
		return
	}

	if err := h.usecase.DeleteInsitutions(id); err != nil {
		helper.InternalServerErrorResponse(c, "failed to delete institution")
		return
	}	
	helper.SuccessResponse(c, "institution deleted successfully", existingInstitution)
}

func (h *InstitutionHandler) GetByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		helper.BadRequestResponse(c, "email query parameter is required")
		return
	}

	institution, err := h.usecase.GetInsitutionByEmail(email)
	if err != nil {
		if errors.Is(err, domain.ErrInstitutionNotFound) {
			helper.NotFoundResponse(c, "institution not found")
			return
		}
		helper.InternalServerErrorResponse(c, "failed to fetch institution by email")
		return
	}

	helper.SuccessResponse(c, "institution retrieved successfully", institution)
}