package helper

import (
	"tindak_ai/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserUUIDFromContext retrieves the user UUID from context safely.
// Returns uuid.Nil if not found or invalid.
func GetUserUUIDFromContext(c *gin.Context) uuid.UUID {
	userIDRaw, exist := c.Get(domain.ContextKeyUserID)
	if !exist {
		return uuid.Nil
	}
	userIDStr, ok := userIDRaw.(string)
	if !ok {
		return uuid.Nil
	}
	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil
	}
	return uid
}
