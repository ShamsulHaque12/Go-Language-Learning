package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"

	"go-language-learning/Week2/Day12/exercise3/models"
	"go-language-learning/Week2/Day12/exercise3/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	for _, prefix := range []string{"/api/users", "/users"} {
		users := r.Group(prefix)
		{
			users.GET("", h.GetAllUsers)
			users.GET("/:id", h.GetUserByID)
			users.POST("", h.CreateUser)
			users.PUT("/:id", h.UpdateUser)
			users.PATCH("/:id", h.PatchUser)
			users.DELETE("/:id", h.DeleteUser)
		}
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
		handleValidationError(c, err)
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
		handleValidationError(c, err)
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
		handleValidationError(c, err)
		return
	}

	if req.Name == nil &&
		req.Email == nil &&
		req.Age == nil &&
		req.IsActive == nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "At least one field is required to update",
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

func handleValidationError(c *gin.Context, err error) {
	details := formatValidationErrors(err)
	if details != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"error":   "Validation failed",
			"details": details,
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error":  "Validation failed: " + err.Error(),
	})
}

func formatValidationErrors(err error) map[string]string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		details := make(map[string]string)
		for _, fe := range ve {
			field := formatFieldName(fe.Field())
			details[field] = formatFieldError(fe)
		}
		return details
	}
	return nil
}

func formatFieldName(field string) string {
	switch field {
	case "Name":
		return "name"
	case "Email":
		return "email"
	case "Age":
		return "age"
	case "IsActive":
		return "is_active"
	default:
		var res strings.Builder
		for i, r := range field {
			if unicode.IsUpper(r) {
				if i > 0 {
					res.WriteRune('_')
				}
				res.WriteRune(unicode.ToLower(r))
			} else {
				res.WriteRune(r)
			}
		}
		return res.String()
	}
}

func formatFieldError(fe validator.FieldError) string {
	field := formatFieldName(fe.Field())
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return "Please provide a valid email address"
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, fe.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, fe.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}