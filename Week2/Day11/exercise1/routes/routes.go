package routes
/*
import (
	"go-language-learning/Week2/Day11/exercise1/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter registers endpoints with injected handlers
func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/users")
	{
		api.GET("", userHandler.GetUsers)
		api.GET("/:id", userHandler.GetUserByID)
		api.POST("", userHandler.CreateUser)
	}

	return r
}
*/