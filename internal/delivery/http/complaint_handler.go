package http

import (
	"encoding/json"
	"errors"
	"fmt" 
	"tindak_ai/internal/domain"
	"tindak_ai/internal/request"
	"tindak_ai/internal/usecase"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ComplaintHandler struct {
	usecase    *usecase.ComplaintUsecase
	authRepo   domain.AuthRepository
	logUsecase *usecase.LoggerUsecase
}

func NewComplaintHandler(router *gin.RouterGroup , uc *usecase.ComplaintUsecase, authRepo domain.AuthRepository, logUc *usecase.LoggerUsecase) {
	handler := &ComplaintHandler{
		usecase: uc,
		authRepo: authRepo,
		logUsecase: logUc,
	}
	complaintGroup := router.Group("/complaints")
	{
		complaintGroup.GET("" ,handler.GetAll)
		complaintGroup.GET("/:id", handler.GetByID)
		complaintGroup.POST("", handler.Create)
		complaintGroup.PUT("/:id", handler.Update)
		complaintGroup.DELETE("/:id",  handler.Delete) 
		complaintGroup.GET("/number/:id", handler.GetByComplaintNumber)

		complaintGroup.GET("/rating/:id",  handler.GetByIDComplaintRating)
		complaintGroup.POST("/rating",  handler.CreateComplaintRating)
		complaintGroup.PUT("/rating/:id",  handler.UpdateComplaintRating)
		complaintGroup.DELETE("/rating/:id",  handler.DeleteComplaintRating)
		// complaintGroup.GET("", middleware.Authorize(authRepo, "read", "institution") ,handler.GetAll)
		// complaintGroup.GET("/:id", middleware.Authorize(authRepo, "show", "institution"), handler.GetByID)
		// complaintGroup.POST("", middleware.Authorize(authRepo, "create", "institution"), handler.Create)
		// complaintGroup.PUT("/:id", middleware.Authorize(authRepo, "update", "institution"), handler.Update)
		// complaintGroup.DELETE("/:id", middleware.Authorize(authRepo, "delete", "institution"), handler.Delete) 

		// complaintGroup.GET("/rating/:id", middleware.Authorize(authRepo, "read", "institution-rating"), handler.GetByIDComplaintRating)
		// complaintGroup.POST("/rating", middleware.Authorize(authRepo, "create", "institution-rating"), handler.CreateComplaintRating)
		// complaintGroup.PUT("/rating/:id", middleware.Authorize(authRepo, "update", "institution-rating"), handler.UpdateComplaintRating)
		// complaintGroup.DELETE("/rating/:id", middleware.Authorize(authRepo, "delete", "institution-rating"), handler.DeleteComplaintRating)
	}
}

func (h *ComplaintHandler) GetAll(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	institution, err := h.usecase.GetAllComplaint()
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch complaints",
			"complaints.GetAllComplaints",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(), 
			helper.PtrString("{}"),
		)
		helper.InternalServerErrorResponse(c, "failed to fetch complaints")
		return
	}
	h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelError,
		"complaints retrieved successfully",
		"complaints.GetAllComplaints", 
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(), 
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "complaints retrieved successfully", institution)
}

func (h *ComplaintHandler) GetByID(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	} 
	institution, err := h.usecase.GetComplaintByID(id)
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch complaint",
			"complaints.GetByIDComplaint",
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
		"complaint retrieved successfully",
		"complaints.GetByIDComplaint",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "complaint retrieved successfully", institution)
}

func (h *ComplaintHandler) GetByComplaintNumber(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")  
	institution, err := h.usecase.GetComplaintByComplaintNumber(idParams)
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch complaint",
			"complaints.GetByComplaintNumber",
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
		"complaint retrieved successfully",
		"complaints.GetByComplaintNumber",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "complaint retrieved successfully", institution)
}


func (h *ComplaintHandler) Create(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	var req request.ComplaintCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}  

	estimatedAt, err := helper.ParseOptionalTimeRFC3339(req.EstimatedAt)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}

	// userID, err := uuid.Parse(req.UserID)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"errors": map[string]string{
	// 			"UserID": "UserID must be a valid UUID",
	// 		},
	// 		"message": "validation error",
	// 		"status": 400,
	// 	})
	// 	return
	// }

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

	complaint := domain.Complaint{
		ID: uuid.New(),  
		// UserID:   userID,
		ComplaintNumber : helper.GenerateComplaintNumber(),
		UserID:   uid,
		InstitutionID: &institutionID,
		Title : req.Title,
		Description : req.Description,
		Location : req.Location,
		Latitude : req.Latitude,
		Longitude : req.Longitude,
		Status : domain.StatusLevel(req.Status), 
		EstimatedAt: estimatedAt,
	}

	file, err := c.FormFile("photo_url")
	if err == nil {
		path := fmt.Sprintf("uploads/complaint/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			meta := map[string]string {
				"created_by": uid.String(), 
			}
			metaStr, _ := json.Marshal(meta)
			h.logUsecase.Log(
				domain.MethodTypePost,
				domain.LogLevelError,
				"failed to save file :" + err.Error(),
				"complaint.Create",
				helper.PtrUUID(uid),
				c.ClientIP(),
				c.Request.UserAgent(),
				helper.PtrString(string(metaStr)),
			)
			helper.InternalServerErrorResponse(c, domain.ErrFailedTosave.Error())
			return
		}
		complaint.PhotoUrl = &path
	} 

	err = h.usecase.CreateComplaint(&complaint)
	if  err != nil { 
		if errors.Is(err, domain.ErrUserNotFound) { 
			helper.NotFoundResponse(c, "user not found")
			return
		}
		if errors.Is(err, domain.ErrInstitutionNotFound) {
			helper.NotFoundResponse(c, "institution not found")
			return
		} 
		meta := map[string]string {
			"created_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePost,
			domain.LogLevelError,
			"failed to create complaint :" + err.Error(),
			"complaint.Create",
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
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"complaint created successfully",
		"complaint.Create",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.CreatedResponse(c, "complaint created successfully", complaint)
}

func (h *ComplaintHandler) Update(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		helper.BadRequestResponse(c, domain.ErrInvalidUUID.Error())
		return
	}

	var req request.ComplaintUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	existingComplaint, err := h.usecase.GetComplaintByID(id)
	if err != nil {
		meta := map[string]string {
			"updated_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to fetch complaint :" + err.Error(),
			"complaint.GetByIDInsitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.NotFoundResponse(c, "complaint not found")
		return
	} 

	if req.Title != nil {
		existingComplaint.Title = *req.Title
	}
	if req.Description != nil {
		existingComplaint.Description = *req.Description
	} 
	if req.Status != nil {
		existingComplaint.Status = domain.StatusLevel(*req.Status)
	} 
	if req.VerifiedBy != nil { 
		parsedUUID, err := uuid.Parse(*req.VerifiedBy)
		if err != nil {
			helper.BadRequestResponse(c, fmt.Sprintf("invalid verified_by UUID: %v", err))
			return
		}
		existingComplaint.VerifiedBy = &parsedUUID
	} 
	if req.StartedAt != nil {
		startedAt, err := helper.ParseOptionalTimeRFC3339(req.StartedAt)
		if err != nil {
			helper.BadRequestResponse(c, "invalid started_at format: "+err.Error())
			return
		}
		existingComplaint.StartedAt = startedAt
	} 
	if req.CompletedAt != nil {
		completedAt, err := helper.ParseOptionalTimeRFC3339(req.CompletedAt)
		if err != nil {
			helper.BadRequestResponse(c, "invalid started_at format: "+err.Error())
			return
		}
		existingComplaint.CompletedAt = completedAt
	} 
	  
	// Optional file upload
	file, err := c.FormFile("logo_url")
	if err == nil {
		path := fmt.Sprintf("uploads/complaint/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			meta := map[string]string {
				"updated_by": uid.String(), 
			}
			metaStr, _ := json.Marshal(meta)
			h.logUsecase.Log(
				domain.MethodTypePut,
				domain.LogLevelError,
				"failed to save file :" + err.Error(),
				"complaint.Update",
				helper.PtrUUID(uid),
				c.ClientIP(),
				c.Request.UserAgent(),
				helper.PtrString(string(metaStr)),
			)
			helper.InternalServerErrorResponse(c, "failed to save file")
			return
		}
		existingComplaint.PhotoUrl = &path // or use &path
	}

	err = h.usecase.UpdateComplaint(existingComplaint)
	if err != nil { 
		if errors.Is(err, domain.ErrComplaintNotFound) {
			helper.NotFoundResponse(c, domain.ErrComplaintNotFound.Error())
			return
		}
		meta := map[string]string {
			"updated_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to update complaint :" + err.Error(),
			"complaint.Update",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to update complaint")
		return
	}
	meta := map[string]string {
		"updated_by": uid.String(), 
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePut,
		domain.LogLevelInfo,
		"complaint updated successfully",
		"complaint.Update",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.SuccessResponse(c, "complaint updated successfully", existingComplaint)
}


func (h *ComplaintHandler) Delete(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}
	existingComplaint, err := h.usecase.GetComplaintByID(id)
	if err != nil {
		meta := map[string]string {
			"deleted_by": uid.String(),
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to fetch complaint",
			"complaint.GetByIDInsitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.NotFoundResponse(c, err.Error())
		return
	}

	if err := h.usecase.DeleteComplaint(id); err != nil {
		meta := map[string]string {
			"deleted_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to delete complaint",
			"complaint.DeleteInsitutions",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to delete complaint")
		return
	}	
	meta := map[string]string {
		"deleted_by": uid.String(), 
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypeDelete,
		domain.LogLevelInfo,
		"complaint deleted successfully",
		"complaint.DeleteInsitutions",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)

	helper.SuccessResponse(c, "complaint deleted successfully", existingComplaint)
}


// complaint rating

func (h *ComplaintHandler) GetByIDComplaintRating(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	} 
	complaintRating, err := h.usecase.GetByIDComplaintRating(id)
	if err != nil {
		h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch complaint rating",
			"complaint-rating.GetByIDComplaintRating",
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
		"complaint rating retrieved successfully",
		"complaint-rating.GetByIDComplaintRating",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "complaint rating retrieved successfully", complaintRating)
}

func (h *ComplaintHandler) CreateComplaintRating(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	var req request.ComplaintRatingCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}   

	// userID, err := uuid.Parse(req.UserID)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"errors": map[string]string{
	// 			"UserID": "UserID must be a valid UUID",
	// 		},
	// 		"message": "validation error",
	// 		"status": 400,
	// 	})
	// 	return
	// } 
	complaintID, err := uuid.Parse(req.ComplaintID)
	if err != nil {
		c.JSON(400, gin.H{
			"errors": map[string]string{
				"ComplaintID": "ComplaintID must be a valid UUID",
			},
			"message": "validation error",
			"status": 400,
   		 }) 
		return
	}

	complaintRating := domain.ComplaintRating{
		ID: uuid.New(),  
		// UserID:   userID,
		UserID:   uid,
		ComplaintID: complaintID,
		Rating: domain.RatingLevel(req.Rating),
		Comment  : req.Comment,
	} 

	err = h.usecase.AddComplaintRating(&complaintRating)
	if  err != nil { 
		if errors.Is(err, domain.ErrUserNotFound) { 
			helper.NotFoundResponse(c, "user not found")
			return
		} 
		if errors.Is(err, domain.ErrComplaintNotFound) { 
			helper.NotFoundResponse(c, "complaint not found")
			return
		} 
		meta := map[string]string {
			"created_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePost,
			domain.LogLevelError,
			"failed to create complaint rating :" + err.Error(),
			"complaint-rating.Create",
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
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"complaint rating created successfully",
		"complaint-rating.Create",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.CreatedResponse(c, "complaint rating created successfully", complaintRating)
}

func (h *ComplaintHandler) UpdateComplaintRating(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		helper.BadRequestResponse(c, domain.ErrInvalidUUID.Error())
		return
	}

	var req request.ComplaintRatingUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	existingComplaintRating, err := h.usecase.GetByIDComplaintRating(id)
	if err != nil {
		meta := map[string]string {
			"updated_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to fetch complaint rating :" + err.Error(),
			"complaint-rating.GetByIDInsitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.NotFoundResponse(c, "complaint rating not found")
		return
	} 
 
	if req.Rating != nil {
		existingComplaintRating.Rating = domain.RatingLevel(*req.Rating)
	} 
	if req.Comment != nil {
		existingComplaintRating.Comment = req.Comment
	}    
	   

	err = h.usecase.UpdateComplaintRating(existingComplaintRating)
	if err != nil { 
		if errors.Is(err, domain.ErrComplaintNotFound) {
			helper.NotFoundResponse(c, domain.ErrComplaintNotFound.Error())
			return
		}
		meta := map[string]string {
			"updated_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to update complaint rating :" + err.Error(),
			"complaint-rating.Update",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to update complaint rating")
		return
	}
	meta := map[string]string {
		"updated_by": uid.String(), 
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePut,
		domain.LogLevelInfo,
		"complaint rating updated successfully",
		"complaint-rating.Update",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.SuccessResponse(c, "complaint rating updated successfully", existingComplaintRating)
}



func (h *ComplaintHandler) DeleteComplaintRating(c *gin.Context) {
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}
	existingComplaintRating, err := h.usecase.GetByIDComplaintRating(id)
	if err != nil {
		meta := map[string]string {
			"deleted_by": uid.String(),
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to fetch complaint rating",
			"complaint-rating.GetByIDComplaint",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.NotFoundResponse(c, err.Error())
		return
	}

	if err := h.usecase.DeleteComplaintRating(id); err != nil {
		meta := map[string]string {
			"deleted_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to delete complaint rating",
			"complaint-rating.DeleteComplaintRating",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to delete complaint rating")
		return
	}	
	meta := map[string]string {
		"deleted_by": uid.String(), 
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypeDelete,
		domain.LogLevelInfo,
		"complaint rating deleted successfully",
		"complaint-rating.DeleteComplaintRating",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)

	helper.SuccessResponse(c, "complaint rating deleted successfully", existingComplaintRating)
}