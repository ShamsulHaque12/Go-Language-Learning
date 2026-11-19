package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
}

type Meta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

// SendSuccess writes a standardized success JSON response
func SendSuccess(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendPaginated writes a standardized paginated response
func SendPaginated(c *gin.Context, statusCode int, message string, data any, page, pageSize int, total int64) {
	totalPages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		totalPages++
	}

	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta: &Meta{
			Page:       page,
			PageSize:   pageSize,
			TotalItems: total,
			TotalPages: totalPages,
		},
	})
}

// SendError writes a standardized error JSON response
func SendError(c *gin.Context, statusCode int, message string, details any) {
	c.JSON(statusCode, APIResponse{
		Success: false,
		Message: message,
		Error:   details,
	})
}

// HandleInternalError handles unexpected system errors gracefully
func HandleInternalError(c *gin.Context, err error) {
	SendError(c, http.StatusInternalServerError, "An unexpected error occurred", err.Error())
}
