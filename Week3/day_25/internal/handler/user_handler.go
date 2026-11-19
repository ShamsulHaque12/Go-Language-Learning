package handler

import (
	"errors"
	"net/http"
	"strconv"

	"day_25/internal/domain"
	"day_25/internal/middleware"
	"day_25/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetProfile returns the current logged-in user profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userIDVal, exists := c.Get(middleware.CtxUserIDKey)
	if !exists {
		SendError(c, http.StatusUnauthorized, "User context not found", nil)
		return
	}

	userID := userIDVal.(uint)
	profile, err := h.userService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			SendError(c, http.StatusNotFound, "User not found", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Failed to retrieve profile", err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "User profile retrieved", profile)
}

// UpdateProfile updates the current logged-in user's profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userIDVal, exists := c.Get(middleware.CtxUserIDKey)
	if !exists {
		SendError(c, http.StatusUnauthorized, "User context not found", nil)
		return
	}

	var req domain.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	userID := userIDVal.(uint)
	updated, err := h.userService.UpdateProfile(c.Request.Context(), userID, &req)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			SendError(c, http.StatusNotFound, "User not found", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Failed to update profile", err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Profile updated successfully", updated)
}

// GetAllUsers returns a paginated list of users (e.g. for admin/listing)
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	users, total, err := h.userService.GetAllUsers(c.Request.Context(), page, pageSize)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}

	SendPaginated(c, http.StatusOK, "Users retrieved successfully", users, page, pageSize, total)
}

// TransferBalance handles funds transfer between user accounts
func (h *UserHandler) TransferBalance(c *gin.Context) {
	fromIDVal, exists := c.Get(middleware.CtxUserIDKey)
	if !exists {
		SendError(c, http.StatusUnauthorized, "User context not found", nil)
		return
	}

	var req domain.TransferBalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "Invalid transfer request", err.Error())
		return
	}

	fromID := fromIDVal.(uint)
	if err := h.userService.TransferBalance(c.Request.Context(), fromID, req.ToUserID, req.Amount); err != nil {
		if errors.Is(err, domain.ErrInsufficientBalance) {
			SendError(c, http.StatusBadRequest, "Transfer rejected", err.Error())
			return
		}
		if errors.Is(err, domain.ErrSameAccountTransfer) || errors.Is(err, domain.ErrInvalidAmount) {
			SendError(c, http.StatusBadRequest, "Invalid transfer parameters", err.Error())
			return
		}
		if errors.Is(err, domain.ErrUserNotFound) {
			SendError(c, http.StatusNotFound, "Account not found", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Transfer execution failed", err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Balance transferred successfully", gin.H{
		"from_user_id": fromID,
		"to_user_id":   req.ToUserID,
		"amount":       req.Amount,
	})
}
