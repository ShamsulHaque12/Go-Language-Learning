package handlers
/*
import (
	"net/http"
	"strconv"

	"go-language-learning/Week2/Day10/mini_project/models"

	"github.com/gin-gonic/gin"
)

// GetUsers handles GET /users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All users fetched successfully",
		"users":   models.GetAllUsers(),
	})
}

// GetUserByID handles GET /users/:id
func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID format",
		})
		return
	}

	user, found := models.GetUserByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}

// CreateUser handles POST /users
func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format",
		})
		return
	}

	if newUser.Name == "" || newUser.Email == "" || newUser.Age <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Name, email, and positive age are required",
		})
		return
	}

	models.CreateUser(&newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    newUser,
	})
}

// UpdateUser handles PUT /users/:id
func UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID format",
		})
		return
	}

	var updateUser models.User
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format",
		})
		return
	}

	if updateUser.Name == "" || updateUser.Email == "" || updateUser.Age <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Name, email, and positive age are required",
		})
		return
	}

	updated, found := models.UpdateUser(id, updateUser)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updated,
	})
}

// DeleteUser handles DELETE /users/:id
func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID format",
		})
		return
	}

	if deleted := models.DeleteUser(id); !deleted {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

*/