package middleware
/*
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIKeyMiddleware verifies the X-API-Key header
func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")

		if apiKey == "" || apiKey != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
*/