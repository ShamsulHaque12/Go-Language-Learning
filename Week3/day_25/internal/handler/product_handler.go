package handler

import (
	"errors"
	"net/http"
	"strconv"

	"day_25/internal/domain"
	"day_25/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct handles POST /api/v1/products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req domain.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "Invalid product payload", err.Error())
		return
	}

	product, err := h.productService.CreateProduct(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateSKU) {
			SendError(c, http.StatusConflict, "Product creation failed", err.Error())
			return
		}
		if errors.Is(err, domain.ErrInvalidPrice) || errors.Is(err, domain.ErrInvalidStock) {
			SendError(c, http.StatusBadRequest, "Validation error", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	SendSuccess(c, http.StatusCreated, "Product created successfully", product)
}

// GetProductByID handles GET /api/v1/products/:id
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		SendError(c, http.StatusBadRequest, "Invalid product ID", "ID must be a positive integer")
		return
	}

	product, err := h.productService.GetProductByID(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			SendError(c, http.StatusNotFound, "Product not found", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Failed to retrieve product", err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Product retrieved successfully", product)
}

// GetAllProducts handles GET /api/v1/products (with filtering & pagination)
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	filter := domain.ProductFilter{
		Category: c.Query("category"),
		Status:   c.Query("status"),
		Search:   c.Query("search"),
		SortBy:   c.DefaultQuery("sort_by", "id"),
		Order:    c.DefaultQuery("order", "asc"),
		Page:     page,
		PageSize: pageSize,
	}

	products, total, err := h.productService.GetAllProducts(c.Request.Context(), filter)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Failed to fetch products", err.Error())
		return
	}

	SendPaginated(c, http.StatusOK, "Products retrieved successfully", products, page, pageSize, total)
}

// UpdateProduct handles PUT /api/v1/products/:id (complete replacement)
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		SendError(c, http.StatusBadRequest, "Invalid product ID", "ID must be a positive integer")
		return
	}

	var req domain.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "Invalid update payload", err.Error())
		return
	}

	updated, err := h.productService.UpdateProduct(c.Request.Context(), uint(id), &req)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			SendError(c, http.StatusNotFound, "Product not found", err.Error())
			return
		}
		if errors.Is(err, domain.ErrDuplicateSKU) {
			SendError(c, http.StatusConflict, "Update failed", err.Error())
			return
		}
		if errors.Is(err, domain.ErrInvalidPrice) || errors.Is(err, domain.ErrInvalidStock) {
			SendError(c, http.StatusBadRequest, "Validation error", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Failed to update product", err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Product updated completely (PUT)", updated)
}

// PatchProduct handles PATCH /api/v1/products/:id (partial update)
func (h *ProductHandler) PatchProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		SendError(c, http.StatusBadRequest, "Invalid product ID", "ID must be a positive integer")
		return
	}

	var req domain.PatchProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "Invalid patch payload", err.Error())
		return
	}

	patched, err := h.productService.PatchProduct(c.Request.Context(), uint(id), &req)
	if err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			SendError(c, http.StatusNotFound, "Product not found", err.Error())
			return
		}
		if errors.Is(err, domain.ErrDuplicateSKU) {
			SendError(c, http.StatusConflict, "Patch failed", err.Error())
			return
		}
		if errors.Is(err, domain.ErrInvalidPrice) || errors.Is(err, domain.ErrInvalidStock) {
			SendError(c, http.StatusBadRequest, "Validation error", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Failed to patch product", err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Product partially updated (PATCH)", patched)
}

// DeleteProduct handles DELETE /api/v1/products/:id
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		SendError(c, http.StatusBadRequest, "Invalid product ID", "ID must be a positive integer")
		return
	}

	if err := h.productService.DeleteProduct(c.Request.Context(), uint(id)); err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			SendError(c, http.StatusNotFound, "Product not found", err.Error())
			return
		}
		SendError(c, http.StatusInternalServerError, "Failed to delete product", err.Error())
		return
	}

	SendSuccess(c, http.StatusOK, "Product deleted successfully (DELETE)", gin.H{"id": id})
}
