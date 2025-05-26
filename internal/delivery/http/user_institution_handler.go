package http

import (
	"encoding/json"
	"errors"
	// "log"
	"tindak_ai/internal/domain"
	"tindak_ai/internal/request"
	"tindak_ai/internal/usecase"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserInstitutionHandler struct {
	usecase *usecase.UserInstitutionUsecase
	authRepo domain.AuthRepository
	logUsecase *usecase.LoggerUsecase
}

func NewUserInstitutionHandler(router *gin.RouterGroup, uc *usecase.UserInstitutionUsecase, authRepo domain.AuthRepository, logUsecase *usecase.LoggerUsecase ){
	handler := &UserInstitutionHandler{
		usecase: uc,
		authRepo: authRepo,
		logUsecase: logUsecase,
	}

	userInstitutionGroup := router.Group("/user-institutions")
	{
		userInstitutionGroup.GET("", handler.GetAllUserInstitutions)
		userInstitutionGroup.GET("/:id", handler.GetUserInstitutionByID)
		userInstitutionGroup.POST("", handler.CreateUserInstitution)
		userInstitutionGroup.PUT("/:id", handler.UpdateUserInstitution)
		userInstitutionGroup.DELETE("/:id", handler.DeleteUserInstitution)
		// userInstitutionGroup.GET("", middleware.Authorize(authRepo, "read", "user_institution"), handler.GetAllUserInstitutions)
		// userInstitutionGroup.GET("/:id", middleware.Authorize(authRepo, "show", "user_institution"),handler.GetUserInstitutionByID)
		// userInstitutionGroup.POST("", middleware.Authorize(authRepo, "create", "user_institution"),handler.CreateUserInstitution)
		// userInstitutionGroup.PUT("/:id", middleware.Authorize(authRepo, "update", "user_institution"),handler.UpdateUserInstitution)
		// userInstitutionGroup.DELETE("/:id", middleware.Authorize(authRepo, "delete", "user_institution"),handler.DeleteUserInstitution)
	}
}

func (h *UserInstitutionHandler) GetAllUserInstitutions(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	userInsitutions, err := h.usecase.GetAll()
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch user institutions",
			"user-institution.GetAllUserInstitutions",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(), 
			helper.PtrString("{}"),
		)
		helper.InternalServerErrorResponse(c, "failed to fetch user institutions")
		return
	}
	h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelInfo,
		"user institutions retrieved successfully",
		"user-institution.GetAllUserInstitutions",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"), 
	)
	helper.SuccessResponse(c, "user institutions retrieved successfully", userInsitutions)
}

func (h *UserInstitutionHandler) GetUserInstitutionByID(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}
	user, err := h.usecase.GetByID(id)
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch user institutions",
			"user-institution.GetUserInstitutionByID",
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
		"user institutions retrieved successfully",
		"user-institution.GetUserInstitutionByID",
		helper.PtrUUID(uid), 
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "user institutions retrieved successfully", user)
}

func (h *UserInstitutionHandler) CreateUserInstitution(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	var req request.UserInstitutionCreateRequest  

	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}  

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(400, gin.H{
			"errors": map[string]string{
				"UserID": "UserID must be a valid UUID",
			},
			"message": "validation error",
			"status": 400,
		})
		return
	}

	institutionID, err := uuid.Parse(req.InstitutionID)
	if err != nil {
		c.JSON(400, gin.H{
			"errors": map[string]string{
				"InstitutionID": "InstitutionID must be a valid UUID",
			},
			"message": "validation error",
			"status": 400,
   		 }) 
		return
	}

	userInsitution := domain.UserInstitution{
		ID:           uuid.New(), 
		UserID:   userID,
		InstitutionID: institutionID,
		Position: req.Position,
	} 
	
	err = h.usecase.Create(&userInsitution)
	if  err != nil {
		// log.Print(err)
		if errors.Is(err, domain.ErrUserNotFound) { 
			helper.NotFoundResponse(c, "user not found")
			return
		}
		if errors.Is(err, domain.ErrInstitutionNotFound) {
			helper.NotFoundResponse(c, "institution not found")
			return
		}
		meta := map[string]string{
			"created_by": uid.String(),
		}
		metaBytes, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePost,
			domain.LogLevelError,
			"failed to create user institution" + err.Error(),
			"user-institution.CreateUserInstitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaBytes)),
		)
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}
	meta := map[string]string{
		"created_by": uid.String(),
		"user_insitution_id":  userInsitution.ID.String(), 
	}
	metaBytes, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"user insitution created successfully",
		"user-institution.CreateUser",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaBytes)),
	)
	helper.CreatedResponse(c, "user insitution created successfully", userInsitution)
}

func (h *UserInstitutionHandler) UpdateUserInstitution(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		helper.BadRequestResponse(c, "invalid UUID")
		return
	}

	var req request.UserInstitutionUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	existingUserInstitution, err := h.usecase.GetByID(id)
	if err != nil {
		meta := map[string]string {
			"updated_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to fetch user institution :" + err.Error(),
			"user-institution.GetUserInsitutionByID",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.NotFoundResponse(c, "user institution not found")
		return
	}

	if req.UserID != nil { 
		userID, err := uuid.Parse(*req.UserID)
		if err != nil {
			c.JSON(400, gin.H{
				"errors": map[string]string{
					"UserID": "UserID must be a valid UUID",
				},
				"message": "validation error",
				"status": 400,
			})
			return
		}
		existingUserInstitution.UserID = userID
	}
	if req.InstitutionID != nil {
		institutionID, err := uuid.Parse(*req.InstitutionID)
		if err != nil {
			c.JSON(400, gin.H{
				"errors": map[string]string{
					"InstitutionID": "InstitutionID must be a valid UUID",
				},
				"message": "validation error",
				"status": 400,
			}) 
			return
		}
		existingUserInstitution.InstitutionID = institutionID
	}  
	if req.Position != nil {
		existingUserInstitution.Position = *req.Position
	}  

	err = h.usecase.Update(existingUserInstitution)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) { 
			helper.NotFoundResponse(c, "user not found")
			return
		}
		if errors.Is(err, domain.ErrInstitutionNotFound) {
			helper.NotFoundResponse(c, "institution not found")
			return
		}
		meta := map[string]string {
			"updated_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to update user institution :" + err.Error(),
			"user-institution.Update",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to update user institution")
		return
	}
	meta := map[string]string {
		"updated_by": uid.String(), 
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePut,
		domain.LogLevelInfo,
		"user institution updated successfully",
		"user-institution.Update",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.SuccessResponse(c, "user institution updated successfully", existingUserInstitution)
}

func (h *UserInstitutionHandler) DeleteUserInstitution(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	userInstitutionId, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, "invalid user instiotution ID")
		return
	}

	existingUser, err := h.usecase.GetByID(userInstitutionId)
	if err != nil {
		meta := map[string]string{
			"deleted_by": uid.String(),
			"user_id":    userInstitutionId.String(),
		}
		metaBytes, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to fetch user institution",
			"user-institution.GetUserInstitutionByID",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaBytes)),
		)
		helper.NotFoundResponse(c, "user institution not found")
		return
	}

	if err := h.usecase.Delete(existingUser.ID); err != nil {
		meta := map[string]string{
			"deleted_by": uid.String(),
			"user_institution_id":    userInstitutionId.String(),
		}
		metaBytes, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to delete user institution", 
			"user-institution.DeleteUserInstitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaBytes)),
		)
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}
	meta := map[string]string{
		"deleted_by": uid.String(),
		"user_institution_id":    userInstitutionId.String(),
	}	
	metaBytes, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypeDelete,
		domain.LogLevelInfo,
		"user institution deleted successfully",
		"user-institution.DeleteUserInstitution",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaBytes)),
	)
	helper.SuccessResponse(c, "user institution deleted successfully", nil)
}
