package handlers
/*
import (
	"net/http"
	"strconv"

	"go-language-learning/Week2/Day11/exercise1/dto"
	"go-language-learning/Week2/Day11/exercise1/services"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests and calls UserService
type UserHandler struct {
	service services.UserService
}

// NewUserHandler initializes UserHandler with UserService dependency
func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.service.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved successfully",
		"data":    users,
	})
}

// GetUserByID handles GET /users/:id
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User retrieved successfully",
		"data":    user,
	})
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	res, err := h.service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    res,
	})
}

*/