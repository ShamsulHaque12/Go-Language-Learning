package middleware

import (
	"net/http"
	"strings"

	"day_25/pkg/token"

	"github.com/gin-gonic/gin"
)

const (
	CtxUserIDKey   = "user_id"
	CtxUserEmailKey = "user_email"
	CtxUserRoleKey  = "user_role"
)

// AuthMiddleware validates JWT bearer tokens in the Authorization header
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Authorization header is required",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Authorization format must be: Bearer <token>",
			})
			return
		}

		claims, err := token.ValidateToken(parts[1], jwtSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		c.Set(CtxUserIDKey, claims.UserID)
		c.Set(CtxUserEmailKey, claims.Email)
		c.Set(CtxUserRoleKey, claims.Role)
		c.Next()
	}
}

// RequireRoles verifies that the authenticated user possesses one of the allowed roles
func RequireRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get(CtxUserRoleKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "forbidden: role not identified",
			})
			return
		}

		userRole, ok := roleVal.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   "forbidden: invalid role format",
			})
			return
		}

		for _, r := range roles {
			if strings.EqualFold(r, userRole) {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "forbidden: insufficient permissions",
		})
	}
}
