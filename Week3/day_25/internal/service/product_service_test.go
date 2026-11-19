package service_test

import (
	"context"
	"errors"
	"testing"

	"day_25/internal/domain"
	"day_25/internal/service"
)

func TestProductService_CRUDOperations(t *testing.T) {
	repo := newMockProductRepository()
	productService := service.NewProductService(repo)
	ctx := context.Background()

	// 1. POST: CreateProduct
	createReq := &domain.CreateProductRequest{
		Name:        "Logitech MX Master 3S",
		Description: "Wireless Performance Mouse",
		SKU:         "MOUSE-MX-001",
		Price:       99.99,
		Stock:       50,
		Category:    "Electronics",
		Status:      "active",
	}

	created, err := productService.CreateProduct(ctx, createReq)
	if err != nil {
		t.Fatalf("failed to create product (POST): %v", err)
	}

	if created.ID == 0 {
		t.Errorf("expected valid product ID, got 0")
	}
	if created.SKU != "MOUSE-MX-001" {
		t.Errorf("expected SKU MOUSE-MX-001, got %s", created.SKU)
	}
	if created.Price != 99.99 {
		t.Errorf("expected price 99.99, got %f", created.Price)
	}

	// 2. GET: GetProductByID
	fetched, err := productService.GetProductByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("failed to get product by ID (GET): %v", err)
	}
	if fetched.Name != "Logitech MX Master 3S" {
		t.Errorf("unexpected product name: %s", fetched.Name)
	}

	// 3. GET: GetAllProducts
	filter := domain.ProductFilter{Page: 1, PageSize: 10}
	all, total, err := productService.GetAllProducts(ctx, filter)
	if err != nil {
		t.Fatalf("failed to get all products (GET): %v", err)
	}
	if total != 1 || len(all) != 1 {
		t.Errorf("expected 1 product, got total=%d len=%d", total, len(all))
	}

	// 4. PUT: UpdateProduct (complete replacement)
	putReq := &domain.UpdateProductRequest{
		Name:        "Logitech MX Master 3S (White Edition)",
		Description: "Premium ergonomic wireless mouse",
		SKU:         "MOUSE-MX-001",
		Price:       109.99,
		Stock:       30,
		Category:    "Electronics",
		Status:      "active",
	}

	updated, err := productService.UpdateProduct(ctx, created.ID, putReq)
	if err != nil {
		t.Fatalf("failed to update product (PUT): %v", err)
	}
	if updated.Name != "Logitech MX Master 3S (White Edition)" {
		t.Errorf("PUT failed: expected updated name, got %s", updated.Name)
	}
	if updated.Price != 109.99 {
		t.Errorf("PUT failed: expected price 109.99, got %f", updated.Price)
	}

	// 5. PATCH: PatchProduct (partial update - updating only price and stock)
	newPrice := 89.99
	newStock := 20
	patchReq := &domain.PatchProductRequest{
		Price: &newPrice,
		Stock: &newStock,
	}

	patched, err := productService.PatchProduct(ctx, created.ID, patchReq)
	if err != nil {
		t.Fatalf("failed to patch product (PATCH): %v", err)
	}
	if patched.Price != 89.99 {
		t.Errorf("PATCH failed: expected price 89.99, got %f", patched.Price)
	}
	if patched.Stock != 20 {
		t.Errorf("PATCH failed: expected stock 20, got %d", patched.Stock)
	}
	// Verify other fields remain untouched
	if patched.Name != "Logitech MX Master 3S (White Edition)" {
		t.Errorf("PATCH corrupted unmodified field Name: %s", patched.Name)
	}

	// 6. DELETE: DeleteProduct
	err = productService.DeleteProduct(ctx, created.ID)
	if err != nil {
		t.Fatalf("failed to delete product (DELETE): %v", err)
	}

	// Verify it's gone
	_, err = productService.GetProductByID(ctx, created.ID)
	if !errors.Is(err, domain.ErrProductNotFound) {
		t.Errorf("expected ErrProductNotFound after DELETE, got %v", err)
	}
}
