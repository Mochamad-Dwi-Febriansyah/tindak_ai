package http

import (
	"encoding/json"
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
	logUsecase *usecase.LoggerUsecase
}

func NewInstitutionHandler(router *gin.RouterGroup , uc *usecase.InstitutionUsecase, authRepo domain.AuthRepository, logUc *usecase.LoggerUsecase) {
	handler := &InstitutionHandler{
		usecase: uc,
		authRepo: authRepo,
		logUsecase: logUc,
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
	uid := helper.GetUserUUIDFromContext(c)
	institution, err := h.usecase.GetAllInsitutions()
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch institutions",
			"institution.GetAllInsitutions",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(), 
			helper.PtrString("{}"),
		)
		helper.InternalServerErrorResponse(c, "failed to fetch institutions")
		return
	}
	h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelError,
		"institutions retrieved successfully",
		"institution.GetAllInsitutions", 
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "institutions retrieved successfully", institution)
}

func (h *InstitutionHandler) GetByID(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	} 
	institution, err := h.usecase.GetByIDInsitution(id)
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch institution",
			"institution.GetByIDInsitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString("{}"),
		)
		helper.NotFoundResponse(c, err.Error())
		return
	} 
	h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelInfo,
		"institution retrieved successfully",
		"institution.GetByIDInsitution",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "institution retrieved successfully", institution)
}

func (h *InstitutionHandler) Create(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
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
			meta := map[string]string {
				"created_by": uid.String(),
				"email": req.ContactEmail,
			}
			metaStr, _ := json.Marshal(meta)
			h.logUsecase.Log(
				domain.MethodTypePost,
				domain.LogLevelError,
				"failed to save file :" + err.Error(),
				"institution.Create",
				helper.PtrUUID(uid),
				c.ClientIP(),
				c.Request.UserAgent(),
				helper.PtrString(string(metaStr)),
			)
			helper.InternalServerErrorResponse(c, domain.ErrFailedTosave.Error())
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
		meta := map[string]string {
			"created_by": uid.String(),
			"email": req.ContactEmail,
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePost,
			domain.LogLevelError,
			"failed to create institution :" + err.Error(),
			"institution.Create",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}
	meta := map[string]string {
		"created_by": uid.String(),
		"email": req.ContactEmail,
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"institution created successfully",
		"institution.Create",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.CreatedResponse(c, "institution created successfully", institution)
}

func (h *InstitutionHandler) Update(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		helper.BadRequestResponse(c, domain.ErrInvalidUUID.Error())
		return
	}

	var req request.InstitutionUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	existingInstitution, err := h.usecase.GetByIDInsitution(id)
	if err != nil {
		meta := map[string]string {
			"updated_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to fetch institution :" + err.Error(),
			"institution.GetByIDInsitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
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
			meta := map[string]string {
				"updated_by": uid.String(),
				"email": *req.ContactEmail,
			}
			metaStr, _ := json.Marshal(meta)
			h.logUsecase.Log(
				domain.MethodTypePut,
				domain.LogLevelError,
				"failed to save file :" + err.Error(),
				"institution.Update",
				helper.PtrUUID(uid),
				c.ClientIP(),
				c.Request.UserAgent(),
				helper.PtrString(string(metaStr)),
			)
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
		meta := map[string]string {
			"updated_by": uid.String(),
			"email": *req.ContactEmail,
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to update institution :" + err.Error(),
			"institution.Update",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to update institution")
		return
	}
	meta := map[string]string {
		"updated_by": uid.String(),
		"email": *req.ContactEmail,
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePut,
		domain.LogLevelInfo,
		"institution updated successfully",
		"institution.Update",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.SuccessResponse(c, "institution updated successfully", existingInstitution)
}

func (h *InstitutionHandler) Delete(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}
	existingInstitution, err := h.usecase.GetByIDInsitution(id)
	if err != nil {
		meta := map[string]string {
			"deleted_by": uid.String(),
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to fetch institution",
			"institution.GetByIDInsitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.NotFoundResponse(c, err.Error())
		return
	}

	if err := h.usecase.DeleteInsitutions(id); err != nil {
		meta := map[string]string {
			"deleted_by": uid.String(),
			"email": existingInstitution.ContactEmail,
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to delete institution",
			"institution.DeleteInsitutions",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to delete institution")
		return
	}	
	meta := map[string]string {
		"deleted_by": uid.String(),
		"email": existingInstitution.ContactEmail,
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypeDelete,
		domain.LogLevelInfo,
		"institution deleted successfully",
		"institution.DeleteInsitutions",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)

	helper.SuccessResponse(c, "institution deleted successfully", existingInstitution)
}

func (h *InstitutionHandler) GetByEmail(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
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
		meta := map[string]string {
			"fetched_by": uid.String(),
			"email": email,
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch institution by email :" + err.Error(),
			"institution.GetInsitutionByEmail",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to fetch institution by email")
		return
	}
	meta := map[string]string {
		"fetched_by": uid.String(),
		"email": email,
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelInfo,
		"institution retrieved successfully by email",
		"institution.GetInsitutionByEmail",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.SuccessResponse(c, "institution retrieved successfully", institution)
}