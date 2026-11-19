package service

import (
	"context"
	"strings"

	"day_25/internal/domain"
	"day_25/internal/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *domain.CreateProductRequest) (*domain.ProductResponse, error)
	GetProductByID(ctx context.Context, id uint) (*domain.ProductResponse, error)
	GetAllProducts(ctx context.Context, filter domain.ProductFilter) ([]domain.ProductResponse, int64, error)
	UpdateProduct(ctx context.Context, id uint, req *domain.UpdateProductRequest) (*domain.ProductResponse, error)
	PatchProduct(ctx context.Context, id uint, req *domain.PatchProductRequest) (*domain.ProductResponse, error)
	DeleteProduct(ctx context.Context, id uint) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *domain.CreateProductRequest) (*domain.ProductResponse, error) {
	if req.Price < 0 {
		return nil, domain.ErrInvalidPrice
	}
	if req.Stock < 0 {
		return nil, domain.ErrInvalidStock
	}

	status := "active"
	if req.Status != "" {
		status = strings.ToLower(req.Status)
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		SKU:         strings.ToUpper(strings.TrimSpace(req.SKU)),
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		Status:      status,
	}

	if err := s.productRepo.Create(ctx, product); err != nil {
		return nil, err
	}

	res := product.ToResponse()
	return &res, nil
}

func (s *productService) GetProductByID(ctx context.Context, id uint) (*domain.ProductResponse, error) {
	product, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	res := product.ToResponse()
	return &res, nil
}

func (s *productService) GetAllProducts(ctx context.Context, filter domain.ProductFilter) ([]domain.ProductResponse, int64, error) {
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 || filter.PageSize > 100 {
		filter.PageSize = 10
	}

	products, total, err := s.productRepo.GetAll(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]domain.ProductResponse, len(products))
	for i, p := range products {
		responses[i] = p.ToResponse()
	}

	return responses, total, nil
}

// UpdateProduct (PUT) completely replaces existing product fields
func (s *productService) UpdateProduct(ctx context.Context, id uint, req *domain.UpdateProductRequest) (*domain.ProductResponse, error) {
	existing, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Price < 0 {
		return nil, domain.ErrInvalidPrice
	}
	if req.Stock < 0 {
		return nil, domain.ErrInvalidStock
	}

	existing.Name = req.Name
	existing.Description = req.Description
	existing.SKU = strings.ToUpper(strings.TrimSpace(req.SKU))
	existing.Price = req.Price
	existing.Stock = req.Stock
	existing.Category = req.Category
	existing.Status = strings.ToLower(req.Status)

	if err := s.productRepo.Update(ctx, existing); err != nil {
		return nil, err
	}

	res := existing.ToResponse()
	return &res, nil
}

// PatchProduct (PATCH) partially modifies non-nil fields
func (s *productService) PatchProduct(ctx context.Context, id uint, req *domain.PatchProductRequest) (*domain.ProductResponse, error) {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.SKU != nil {
		updates["sku"] = strings.ToUpper(strings.TrimSpace(*req.SKU))
	}
	if req.Price != nil {
		if *req.Price < 0 {
			return nil, domain.ErrInvalidPrice
		}
		updates["price"] = *req.Price
	}
	if req.Stock != nil {
		if *req.Stock < 0 {
			return nil, domain.ErrInvalidStock
		}
		updates["stock"] = *req.Stock
	}
	if req.Category != nil {
		updates["category"] = *req.Category
	}
	if req.Status != nil {
		updates["status"] = strings.ToLower(*req.Status)
	}

	if len(updates) == 0 {
		// Nothing to update, return current state
		return s.GetProductByID(ctx, id)
	}

	updated, err := s.productRepo.Patch(ctx, id, updates)
	if err != nil {
		return nil, err
	}

	res := updated.ToResponse()
	return &res, nil
}

func (s *productService) DeleteProduct(ctx context.Context, id uint) error {
	return s.productRepo.Delete(ctx, id)
}
