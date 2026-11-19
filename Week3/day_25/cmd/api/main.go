package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"day_25/config"
	"day_25/internal/handler"
	"day_25/internal/repository"
	"day_25/internal/service"
	"day_25/pkg/database"
)

func main() {
	log.Println("[INFO] Starting Product Management REST API Server...")

	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Database Connection (PostgreSQL + GORM)
	db, err := database.ConnectPostgres(
		cfg.DatabaseURL,
		cfg.DBMaxOpenConns,
		cfg.DBMaxIdleConns,
		cfg.DBConnMaxLifetime,
		cfg.Env == "production",
	)
	if err != nil {
		log.Fatalf("[FATAL] Failed to connect database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("[ERROR] Failed to close database: %v", err)
		}
	}()

	// 3. Initialize Repositories & Auto-Migrate Schemas
	userRepo := repository.NewUserRepository(db.DB)
	if err := userRepo.Migrate(); err != nil {
		log.Fatalf("[FATAL] User schema migration failed: %v", err)
	}

	productRepo := repository.NewProductRepository(db.DB)
	if err := productRepo.Migrate(); err != nil {
		log.Fatalf("[FATAL] Product schema migration failed: %v", err)
	}
	log.Println("[INFO] Database schema migration completed successfully (Users & Products)")

	// 4. Initialize Services (Business Logic)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiresIn)
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)

	// 5. Initialize Handlers (Delivery / Presentation)
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	// 6. Setup Gin Router
	router := handler.SetupRouter(handler.RouterConfig{
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		ProductHandler: productHandler,
		JWTSecret:      cfg.JWTSecret,
	})

	// 7. Configure HTTP Server
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 8. Run Server in a Goroutine
	go func() {
		log.Printf("[INFO] Server listening on http://localhost%s", serverAddr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[FATAL] Server listen error: %v", err)
		}
	}()

	// 9. Graceful Shutdown on SIGINT (Ctrl+C) / SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[INFO] Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[FATAL] Server forced to shutdown: %v", err)
	}

	log.Println("[INFO] Server exiting properly. Goodbye!")
}
