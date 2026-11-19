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
		// Extract URL Parameter :id
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

		// User not found
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	})

	// 3. POST /users -> Create a new user
	r.POST("/users", func(c *gin.Context) {
		var newUser User

		// Decode JSON body into newUser struct
		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid JSON format",
			})
			return
		}

		// Required fields validation
		if newUser.Name == "" || newUser.Email == "" || newUser.Age <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Name, email, and positive age are required",
			})
			return
		}

		// Generate auto-increment ID
		maxID := 0
		for _, u := range users {
			if u.ID > maxID {
				maxID = u.ID
			}
		}
		newUser.ID = maxID + 1

		// Append new user to in-memory slice
		users = append(users, newUser)

		// Return HTTP 201 Created with standard response structure
		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    newUser,
		})
	})

	fmt.Println("Server starting on port :8080...")
	// Start server on localhost:8080
	r.Run(":8080")
}

*/
