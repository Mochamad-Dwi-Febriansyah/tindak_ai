package middleware

import (
	"tindak_ai/internal/domain"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Authorize(repo domain.AuthRepository, action string, resource string) gin.HandlerFunc {
	return func(c *gin.Context){
		userIDVal, exist := c.Get(domain.ContextKeyUserID)
		if !exist {
			helper.UnauthorizedResponse(c, "User ID not found in context")
			c.Abort()
			return
		}

		userID, ok := userIDVal.(string)
		if !ok {
			helper.UnauthorizedResponse(c, "Invalid User ID type")
			c.Abort()
			return
		}

		userIDUID, err := uuid.Parse(userID)
		if err != nil {
			helper.UnauthorizedResponse(c, "Invalid UUID format")
			c.Abort()
			return
		}

		allowed, err := repo.HasPermission(userIDUID, action, resource)
		if err != nil {
			helper.InternalServerErrorResponse(c, "Failed to check permission")
			c.Abort()
			return
		}

		if !allowed {
			helper.ForbiddenResponse(c, "You do not have permission to access this resource")
			c.Abort()
			return
		}

		c.Next()
	}
}