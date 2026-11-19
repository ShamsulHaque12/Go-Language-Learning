package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"go-language-learning/Week2/Day12/exercise3/config"
	"go-language-learning/Week2/Day12/exercise3/database"
	"go-language-learning/Week2/Day12/exercise3/handler"
	"go-language-learning/Week2/Day12/exercise3/repository"
	"go-language-learning/Week2/Day12/exercise3/service"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	// Ensure users table exists
	if err := database.CreateUserTable(db); err != nil {
		log.Fatalf("failed to create users table: %v", err)
	}

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepo)

	// Initialize handler
	userHandler := handler.NewUserHandler(userService)

	// Setup Gin router
	r := gin.Default()

	// User routes (accessible at /api/users and /users)
	for _, prefix := range []string{"/api/users", "/users"} {
		api := r.Group(prefix)
		{
			api.GET("", userHandler.GetAllUsers)
			api.GET("/:id", userHandler.GetUserByID)
			api.POST("", userHandler.CreateUser)
			api.PUT("/:id", userHandler.UpdateUser)
			api.PATCH("/:id", userHandler.PatchUser)
			api.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	// Start server
	if err := r.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
