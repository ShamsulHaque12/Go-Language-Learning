package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"day_25/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	GetByID(ctx context.Context, id uint) (*domain.Product, error)
	GetBySKU(ctx context.Context, sku string) (*domain.Product, error)
	GetAll(ctx context.Context, filter domain.ProductFilter) ([]domain.Product, int64, error)
	Update(ctx context.Context, product *domain.Product) error
	Patch(ctx context.Context, id uint, updates map[string]interface{}) (*domain.Product, error)
	Delete(ctx context.Context, id uint) error
	Migrate() error
}

type postgresProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &postgresProductRepository{db: db}
}

// Migrate auto-migrates the products table
func (r *postgresProductRepository) Migrate() error {
	return r.db.AutoMigrate(&domain.Product{})
}

// Create inserts a new product record
func (r *postgresProductRepository) Create(ctx context.Context, product *domain.Product) error {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "UNIQUE") {
			return domain.ErrDuplicateSKU
		}
		return err
	}
	return nil
}

// GetByID retrieves a product by its ID
func (r *postgresProductRepository) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	var product domain.Product
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &product, nil
}

// GetBySKU retrieves a product by its unique SKU
func (r *postgresProductRepository) GetBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	var product domain.Product
	if err := r.db.WithContext(ctx).Where("sku = ?", sku).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &product, nil
}

// GetAll returns a filtered and paginated list of products
func (r *postgresProductRepository) GetAll(ctx context.Context, filter domain.ProductFilter) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Product{})

	// Filter by Category
	if filter.Category != "" {
		query = query.Where("LOWER(category) = LOWER(?)", filter.Category)
	}

	// Filter by Status
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	// Search by name or description or SKU
	if filter.Search != "" {
		searchTerm := "%" + strings.ToLower(filter.Search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ? OR LOWER(sku) LIKE ?", searchTerm, searchTerm, searchTerm)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting
	sortBy := "id"
	if filter.SortBy == "price" || filter.SortBy == "name" || filter.SortBy == "created_at" || filter.SortBy == "stock" {
		sortBy = filter.SortBy
	}

	order := "ASC"
	if strings.ToUpper(filter.Order) == "DESC" {
		order = "DESC"
	}

	offset := (filter.Page - 1) * filter.PageSize
	if err := query.Order(fmt.Sprintf("%s %s", sortBy, order)).Offset(offset).Limit(filter.PageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// Update completely replaces a product record (PUT)
func (r *postgresProductRepository) Update(ctx context.Context, product *domain.Product) error {
	result := r.db.WithContext(ctx).Save(product)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") || strings.Contains(result.Error.Error(), "UNIQUE") {
			return domain.ErrDuplicateSKU
		}
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}

// Patch updates specific non-nil fields of a product (PATCH)
func (r *postgresProductRepository) Patch(ctx context.Context, id uint, updates map[string]interface{}) (*domain.Product, error) {
	var product domain.Product

	// Check if product exists
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}

	// Perform partial update
	result := r.db.WithContext(ctx).Model(&product).Updates(updates)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") || strings.Contains(result.Error.Error(), "UNIQUE") {
			return nil, domain.ErrDuplicateSKU
		}
		return nil, result.Error
	}

	// Re-fetch updated product
	if err := r.db.WithContext(ctx).First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

// Delete removes a product record (DELETE)
func (r *postgresProductRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&domain.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrProductNotFound
	}
	return nil
}
