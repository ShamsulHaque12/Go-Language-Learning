package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"go-language-learning/Week2/Day12/exercise2/models"
	"go-language-learning/Week2/Day12/exercise2/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("", h.GetAllUsers)
		users.GET("/:id", h.GetUserByID)
		users.POST("", h.CreateUser)
		users.PUT("/:id", h.UpdateUser)
		users.PATCH("/:id", h.PatchUser)
		users.DELETE("/:id", h.DeleteUser)
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Validation failed: " + err.Error(),
		})
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Validation failed: " + err.Error(),
		})
		return
	}

	user, err := h.service.UpdateUser(id, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) PatchUser(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	var req models.PatchUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Validation failed: " + err.Error(),
		})
		return
	}

	user, err := h.service.PatchUser(id, &req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User deleted successfully",
	})
}

func parseID(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "User ID must be a positive integer",
		})
		return 0, errors.New("invalid ID")
	}
	return id, nil
}

func (h *UserHandler) handleError(c *gin.Context, err error) {
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  "User not found",
		})
		return
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		c.JSON(http.StatusConflict, gin.H{
			"status": "error",
			"error":  "Email already exists",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"status": "error",
		"error":  "Internal server error",
	})
}
