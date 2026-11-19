package handler

import (
	"errors"
	"net/http"

	"day_25/internal/domain"
	"day_25/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body domain.RegisterRequest true "User Registration"
// @Success 201 {object} APIResponse
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	res, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyExists) {
			SendError(c, http.StatusConflict, "Registration failed", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Registration failed", err.Error())
		return
	}

	SendSuccess(c, http.StatusCreated, "User registered successfully", res)
}

// Login handles user authentication
// @Summary Authenticate user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body domain.LoginRequest true "User Login"
// @Success 200 {object} APIResponse
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	res, err := h.authService.Login(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			SendError(c, http.StatusUnauthorized, "Authentication failed", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Login failed", err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Login successful", res)
}
