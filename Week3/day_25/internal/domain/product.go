package domain

import (
	"time"
)

// Product represents the core domain model and database entity
type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(150);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	SKU         string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"sku"`
	Price       float64   `gorm:"type:decimal(12,2);not null;default:0.00" json:"price"`
	Stock       int       `gorm:"not null;default:0" json:"stock"`
	Category    string    `gorm:"type:varchar(100);index;not null" json:"category"`
	Status      string    `gorm:"type:varchar(50);default:'active'" json:"status"` // active, inactive, draft
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// CreateProductRequest (POST) payload
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=2,max=150"`
	Description string  `json:"description"`
	SKU         string  `json:"sku" binding:"required,min=2,max=100"`
	Price       float64 `json:"price" binding:"required,gte=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
	Category    string  `json:"category" binding:"required,min=2,max=100"`
	Status      string  `json:"status"`
}

// UpdateProductRequest (PUT - complete replacement) payload
type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=2,max=150"`
	Description string  `json:"description"`
	SKU         string  `json:"sku" binding:"required,min=2,max=100"`
	Price       float64 `json:"price" binding:"required,gte=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
	Category    string  `json:"category" binding:"required,min=2,max=100"`
	Status      string  `json:"status" binding:"required,oneof=active inactive draft"`
}

// PatchProductRequest (PATCH - partial update) payload with pointer fields
type PatchProductRequest struct {
	Name        *string  `json:"name,omitempty" binding:"omitempty,min=2,max=150"`
	Description *string  `json:"description,omitempty"`
	SKU         *string  `json:"sku,omitempty" binding:"omitempty,min=2,max=100"`
	Price       *float64 `json:"price,omitempty" binding:"omitempty,gte=0"`
	Stock       *int     `json:"stock,omitempty" binding:"omitempty,gte=0"`
	Category    *string  `json:"category,omitempty" binding:"omitempty,min=2,max=100"`
	Status      *string  `json:"status,omitempty" binding:"omitempty,oneof=active inactive draft"`
}

// ProductFilter parameters for GET /products queries
type ProductFilter struct {
	Category string
	Status   string
	Search   string
	Page     int
	PageSize int
	SortBy   string
	Order    string
}

// ProductResponse represents client-safe response DTO
type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SKU         string    `json:"sku"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) ToResponse() ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		SKU:         p.SKU,
		Price:       p.Price,
		Stock:       p.Stock,
		Category:    p.Category,
		Status:      p.Status,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
