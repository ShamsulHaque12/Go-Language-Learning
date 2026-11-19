package main
/*
import (
	"fmt"

	"go-language-learning/Week2/Day11/exercise1/handlers"
	"go-language-learning/Week2/Day11/exercise1/repositories"
	"go-language-learning/Week2/Day11/exercise1/routes"
	"go-language-learning/Week2/Day11/exercise1/services"
)

func main() {
	// 1. Initialize Data Access Layer (Repository)
	userRepo := repositories.NewUserRepository()

	// 2. Inject Repository into Business Logic Layer (Service)
	userService := services.NewUserService(userRepo)

	// 3. Inject Service into Transport Layer (Handler)
	userHandler := handlers.NewUserHandler(userService)

	// 4. Setup Router with Handlers
	r := routes.SetupRouter(userHandler)

	fmt.Println("Starting Clean Architecture Server on port :8080...")
	// 5. Run Server
	r.Run(":8080")
}
*/