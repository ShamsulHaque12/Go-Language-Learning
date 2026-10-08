package routes
/*
import (
	"go-language-learning/Week2/Day10/mini_project/handlers"
	"go-language-learning/Week2/Day10/mini_project/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes Gin router, registers middlewares and routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global Middleware
	r.Use(middleware.LoggerMiddleware())

	// Route Group protected by APIKeyMiddleware
	api := r.Group("/users")
	api.Use(middleware.APIKeyMiddleware())
	{
		api.GET("", handlers.GetUsers)
		api.GET("/:id", handlers.GetUserByID)
		api.POST("", handlers.CreateUser)
		api.PUT("/:id", handlers.UpdateUser)
		api.DELETE("/:id", handlers.DeleteUser)
	}

	return r
}
*/