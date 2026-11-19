package handler

import (
	"net/http"

	"day_25/internal/middleware"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	AuthHandler    *AuthHandler
	UserHandler    *UserHandler
	ProductHandler *ProductHandler
	JWTSecret      string
}

func SetupRouter(cfg RouterConfig) *gin.Engine {
	r := gin.New()

	// Global Middlewares
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestLogger())
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "Go REST API Clean Architecture - Product Management",
		})
	})

	// About endpoint
	r.GET("/about", aboutHandler)

	// Root Welcome & API Directory
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Product Management REST API (Clean Architecture)",
			"status":  "running",
			"health":  "http://localhost:8080/health",
			"endpoints": gin.H{
				"products": gin.H{
					"list_all": "GET /api/v1/products (supports ?category=&search=&page=&page_size=)",
					"get_by_id": "GET /api/v1/products/:id",
					"create":   "POST /api/v1/products",
					"update":   "PUT /api/v1/products/:id (full replace)",
					"patch":    "PATCH /api/v1/products/:id (partial update)",
					"delete":   "DELETE /api/v1/products/:id",
				},
				"auth": gin.H{
					"register": "POST /api/v1/auth/register",
					"login":    "POST /api/v1/auth/login",
				},
				"users": gin.H{
					"profile":  "GET /api/v1/users/me (Protected)",
					"transfer": "POST /api/v1/users/transfer (Protected)",
				},
			},
		})
	})

	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", cfg.AuthHandler.Register)
			auth.POST("/login", cfg.AuthHandler.Login)
		}

		// User routes (Protected)
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			users.GET("/me", cfg.UserHandler.GetProfile)
			users.PUT("/me", cfg.UserHandler.UpdateProfile)
			users.POST("/transfer", cfg.UserHandler.TransferBalance)
			users.GET("", cfg.UserHandler.GetAllUsers)
		}

		// Product Management routes (GET, POST, PUT, PATCH, DELETE)
		products := api.Group("/products")
		{
			// Read operations
			products.GET("", cfg.ProductHandler.GetAllProducts)
			products.GET("/:id", cfg.ProductHandler.GetProductByID)

			// Write / Mutation operations (GET, PUT, POST, PATCH, DELETE)
			products.POST("", cfg.ProductHandler.CreateProduct)
			products.PUT("/:id", cfg.ProductHandler.UpdateProduct)
			products.PATCH("/:id", cfg.ProductHandler.PatchProduct)
			products.DELETE("/:id", cfg.ProductHandler.DeleteProduct)
		}
	}

	return r
}

// aboutHandler returns application metadata and information
func aboutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"app_name":    "Product Management REST API",
		"version":     "1.0.0",
		"developer":   "Shamsul Haque Sujon",
		"description": "A production-grade RESTful API built with Go, Gin, GORM, and Clean Architecture.",
		"technology": []string{
			"Go 1.23",
			"Gin Web Framework",
			"PostgreSQL (Neon Cloud)",
			"GORM ORM",
			"JWT Authentication",
			"Bcrypt Cryptography",
		},
		"supported_methods": []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		"status":            "running",
	})
}
