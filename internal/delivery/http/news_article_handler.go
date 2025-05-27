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
	"github.com/gosimple/slug"
)

type NewsArticleHandler struct {
	usecase *usecase.NewsArticleUsecase
	authRepo domain.AuthRepository
	logUsecase *usecase.LoggerUsecase
}

func NewNewsArticleHandler(router *gin.RouterGroup, uc *usecase.NewsArticleUsecase, authRepo domain.AuthRepository, logUsecase *usecase.LoggerUsecase){
	handler := &NewsArticleHandler{
		usecase: uc,
		authRepo: authRepo,
		logUsecase: logUsecase,
	}

	newsArticleGroup := router.Group("/news-article")
	{
		newsArticleGroup.GET("", handler.GetAllNewsArticle)
		newsArticleGroup.GET("/:id",  handler.GetNewsArticleByID)
		newsArticleGroup.POST("", handler.CreateNewsArticle) 
		newsArticleGroup.PUT("/:id", handler.UpdateNewsArticle)
		newsArticleGroup.DELETE("/:id",  handler.DeleteNewsArticle) 
		// newsArticleGroup.GET("", middleware.Authorize(authRepo, "read", "news-article") ,handler.GetAllNewsArticle)
		// newsArticleGroup.GET("/:id", middleware.Authorize(authRepo, "show", "news-article"), handler.GetNewsArticleByID)
		// newsArticleGroup.POST("", middleware.Authorize(authRepo, "create", "news-article"), handler.CreateNewsArticle)
		// newsArticleGroup.PUT("/:id", middleware.Authorize(authRepo, "update", "news-article"), handler.UpdateNewsArticle)
		// newsArticleGroup.DELETE("/:id", middleware.Authorize(authRepo, "delete", "news-article"), handler.DeleteNewsArticle)  
	}
}

func (h *NewsArticleHandler) GetAllNewsArticle(c *gin.Context){ 
	uid := helper.GetUserUUIDFromContext(c)
	slug := c.Query("slug")
	filter := domain.NewsArticleFilter{
		Slug:  slug, 
	}
	newsArticles, err := h.usecase.GetAllNewsArticle(filter)
	if  err != nil {
			h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"failed to fetch news article : " + err.Error(),
			"news-article.GetAllNewsArticle",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(), 
			helper.PtrString("{}"),
		)
		helper.InternalServerErrorResponse(c, "failed to fetch news article")
		return
	}
	h.logUsecase.Log(
			domain.MethodTypeGet,
			domain.LogLevelError,
			"news article retrieved successfully",
			"news-article.GetAllNewsArticle",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(), 
			helper.PtrString("{}"),
		)
	helper.SuccessResponse(c, "news article retreived successfully", newsArticles)
}

func (h *NewsArticleHandler) GetNewsArticleByID(c *gin.Context){
	uid := helper.GetUserUUIDFromContext(c)
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return 
	}
	newsArticle , err := h.usecase.GetNewsArticleByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrNewsArticleNotFound) {
			h.logUsecase.Log(
				domain.MethodTypeGet,
				domain.LogLevelError,
				"failed to fetch news article :" + err.Error(),
				"news-article.GetByIDNewsArticle",
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
				domain.LogLevelError,
				"failed to fetch news article :" + err.Error(),
				"news-article.GetByIDNewsArticle",
				helper.PtrUUID(uid),
				c.ClientIP(),
				c.Request.UserAgent(),
				helper.PtrString("{}"),
			)
		helper.InternalServerErrorResponse(c, err.Error())
		return 
	}
	h.logUsecase.Log(
		domain.MethodTypeGet,
		domain.LogLevelInfo,
		"news article retrieved successfully",
		"news-article.GetByIDNewsArticle",
		helper.PtrUUID(uid), 
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString("{}"),
	)
	helper.SuccessResponse(c, "news article retrived successfully", newsArticle)
}

func (h *NewsArticleHandler) CreateNewsArticle(c *gin.Context){
	uid := helper.GetUserUUIDFromContext(c)
	var req request.NewsArticleCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}
	
	generatedSlug := slug.Make(req.Title)

	newsArticle := domain.NewsArticle{
		ID : uuid.New(),
		Slug : generatedSlug,
		Title : req.Title,
		Content : req.Content,
		AuthorID : uid,
		IsPublished : req.IsPublished,
		EstimatedAt: req.EstimatedAt,
	}
	file, err := c.FormFile("thumbnail_url")
	if err  == nil {
		path := fmt.Sprintf("uploads/news-article/thumbnail/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			meta := map[string]string{
				"created_by" : uid.String(), 
			}
			metaStr, _ := json.Marshal(meta)
			h.logUsecase.Log(
				domain.MethodTypePost,
				domain.LogLevelError,
				"failed to save file :" + err.Error(),
				"news-article.Create",
				helper.PtrUUID(uid),
				c.ClientIP(),
				c.Request.UserAgent(),
				helper.PtrString(string(metaStr)),
			)
			helper.InternalServerErrorResponse(c, domain.ErrFailedTosave.Error())
			return
		}
		newsArticle.ThumbnailUrl = &path
	}
	err = h.usecase.CreateNewsArticle(&newsArticle)
	if err != nil {
		meta := map[string]string {
			"created_by" : uid.String(),
		}
		metaStr , _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePost,
			domain.LogLevelError,
			"failed to create news article :" + err.Error(),
			"news-article.Create",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, err.Error())
		return
	}
	meta := map[string]string{
		"created_by" : uid.String(),
	}
	metaStr , _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePost,
		domain.LogLevelInfo,
		"news article created successfully",
		"news-article.Create",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.SuccessResponse(c, "news article created successfully", newsArticle) 
}

func (h *NewsArticleHandler) UpdateNewsArticle(c *gin.Context){
	uid := helper.GetUserUUIDFromContext(c)
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil{
		helper.BadRequestResponse(c, domain.ErrInvalidUUID.Error())
		return
	}

	var req request.NewsArticleUpdateRequest
	if err := c.ShouldBind(&req); err != nil {
		helper.ValidationErrorResponse(c, err)
		return
	}

	existingNewsArticle, err := h.usecase.GetNewsArticleByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrNewsArticleNotFound) {
			h.logUsecase.Log(
				domain.MethodTypePut,
				domain.LogLevelError,
				"failed to fetch news article :" + err.Error(),
				"news-article.GetByIDNewsArticle",
				helper.PtrUUID(uid),
				c.ClientIP(),
				c.Request.UserAgent(),
				helper.PtrString("{}"),
			)
			helper.NotFoundResponse(c, err.Error())
			return
		}
		helper.InternalServerErrorResponse(c, err.Error())
		return 
	}

	if req.Title != nil {
		existingNewsArticle.Title = *req.Title
		generatedSlug := slug.Make(*req.Title)
		existingNewsArticle.Slug = generatedSlug
	}
	if req.Content != nil {
		existingNewsArticle.Content = *req.Content
	} 

	file, err := c.FormFile("thumbnail_url")
	if err == nil {
		path := fmt.Sprintf("uploads/news-article/thumbnail/%s", file.Filename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			meta := map[string]string {
				"updated_by": uid.String(), 
			}
			metaStr, _ := json.Marshal(meta)
			h.logUsecase.Log(
				domain.MethodTypePut,
				domain.LogLevelError,
				"failed to save file :" + err.Error(),
				"news-article.Update",
				helper.PtrUUID(uid),
				c.ClientIP(),
				c.Request.UserAgent(),
				helper.PtrString(string(metaStr)),
			)
			helper.InternalServerErrorResponse(c, domain.ErrFailedTosave.Error())
			return
		}
		existingNewsArticle.ThumbnailUrl = &path // or use &path
	}
	err = h.usecase.UpdateNewsArticle(existingNewsArticle)
	if err != nil { 
		if errors.Is(err, domain.ErrNewsArticleNotFound) {
			helper.NotFoundResponse(c, domain.ErrNewsArticleNotFound.Error())
			return
		}
		meta := map[string]string {
			"updated_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypePut,
			domain.LogLevelError,
			"failed to update news article :" + err.Error(),
			"news-article.Update",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to update news article")
		return
	}
	meta := map[string]string {
		"updated_by": uid.String(), 
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypePut,
		domain.LogLevelInfo,
		"news article updated successfully",
		"news-article.Update",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)
	helper.SuccessResponse(c, "news article updated successfully", existingNewsArticle)
}

func (h *NewsArticleHandler) DeleteNewsArticle(c *gin.Context){
	uid := helper.GetUserUUIDFromContext(c)
	idParams := c.Param("id")
	id, err := uuid.Parse(idParams)
	if err != nil {
		helper.BadRequestResponse(c, err.Error())
		return
	}
	existingNewsArticle, err := h.usecase.GetNewsArticleByID(id)
	if err != nil {
		meta := map[string]string {
			"deleted_by": uid.String(),
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to fetch news article",
			"news-article.GetByIDInsitution",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.NotFoundResponse(c, err.Error())
		return
	}

	if err := h.usecase.DeleteNewsArticle(id); err != nil {
		meta := map[string]string {
			"deleted_by": uid.String(), 
		}
		metaStr, _ := json.Marshal(meta)
		h.logUsecase.Log(
			domain.MethodTypeDelete,
			domain.LogLevelError,
			"failed to delete news article : "+ err.Error(),
			"news-article.DeleteInsitutions",
			helper.PtrUUID(uid),
			c.ClientIP(),
			c.Request.UserAgent(),
			helper.PtrString(string(metaStr)),
		)
		helper.InternalServerErrorResponse(c, "failed to delete news article")
		return
	}	
	meta := map[string]string {
		"deleted_by": uid.String(), 
	}
	metaStr, _ := json.Marshal(meta)
	h.logUsecase.Log(
		domain.MethodTypeDelete,
		domain.LogLevelInfo,
		"news article deleted successfully",
		"news-article.DeleteInsitutions",
		helper.PtrUUID(uid),
		c.ClientIP(),
		c.Request.UserAgent(),
		helper.PtrString(string(metaStr)),
	)

	helper.SuccessResponse(c, "news article deleted successfully", existingNewsArticle)
} 