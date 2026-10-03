package main

/*
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// In-memory sample database
var users = []User{
	{ID: 1, Name: "Sujon", Email: "sujon@gmail.com", Age: 27},
	{ID: 2, Name: "Rahim", Email: "rahim@gmail.com", Age: 25},
	{ID: 3, Name: "Karim", Email: "karim@gmail.com", Age: 30},
}

func main() {
	// Create a default Gin router (includes Logger & Recovery middleware)
	r := gin.Default()

	// 1. GET /users -> Return all users
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "All users fetched successfully",
			"users":   users,
		})
	})

	// 2. GET /users/:id -> Return single user by ID
	r.GET("/users/:id", func(c *gin.Context) {
		// URL Parameter :id extract করা
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid user ID format",
			})
			return
		}

		// Search user by ID
		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, gin.H{
					"message": "User fetched successfully",
					"user":    user,
				})
				return
			}
		}

		// User না পাওয়া গেলে 404 Not Found
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	})

	fmt.Println("Server starting on port :8080...")
	// Start server on localhost:8080
	r.Run(":8080")
}

*/