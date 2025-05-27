package middleware

import (
	// "strings"
	"tindak_ai/internal/domain"
	"tindak_ai/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// authHeader := c.GetHeader("Authorization")
		cookie, err:= c.Cookie("jwt_token")
		if err != nil {
			helper.UnauthorizedResponse(c, "Authorization header missing")
			c.Abort()
			return
		}

		// tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenMalformed
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			helper.UnauthorizedResponse(c, "Invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			helper.UnauthorizedResponse(c, "Invalid token claims")
			c.Abort()
			return
		}
		c.Set(domain.ContextKeyUserID, claims["user_id"]) 
		c.Next()
	}
}