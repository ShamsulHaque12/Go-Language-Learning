package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"go-language-learning/Week2/Day12/exercise2/config"
	"go-language-learning/Week2/Day12/exercise2/database"
	"go-language-learning/Week2/Day12/exercise2/handler"
	"go-language-learning/Week2/Day12/exercise2/repository"
	"go-language-learning/Week2/Day12/exercise2/service"
)

func main() {
	// 1. Load configuration (.env)
	cfg := config.LoadConfig()

	// 2. Database connection with connection pool
	db, err := database.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// 3. Initialize table
	if err := database.CreateUsersTable(db); err != nil {
		log.Fatalf("Table creation failed: %v", err)
	}

	// 4. Dependency Injection (Repository -> Service -> Handler)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 5. Setup Gin Router
	r := gin.Default()
	userHandler.RegisterRoutes(r)

	// 6. Start Server
	log.Printf("Starting server on port :%s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}