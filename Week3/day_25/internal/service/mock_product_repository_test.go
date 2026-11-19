package service_test

import (
	"context"
	"strings"

	"day_25/internal/domain"
	"day_25/internal/repository"
)

type mockProductRepository struct {
	products map[uint]*domain.Product
	bySKU    map[string]*domain.Product
	nextID   uint
}

func newMockProductRepository() repository.ProductRepository {
	return &mockProductRepository{
		products: make(map[uint]*domain.Product),
		bySKU:    make(map[string]*domain.Product),
		nextID:   1,
	}
}

func (m *mockProductRepository) Migrate() error {
	return nil
}

func (m *mockProductRepository) Create(ctx context.Context, product *domain.Product) error {
	sku := strings.ToUpper(product.SKU)
	if _, exists := m.bySKU[sku]; exists {
		return domain.ErrDuplicateSKU
	}

	product.ID = m.nextID
	m.nextID++
	m.products[product.ID] = product
	m.bySKU[sku] = product
	return nil
}

func (m *mockProductRepository) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	p, exists := m.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}
	return p, nil
}

func (m *mockProductRepository) GetBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	p, exists := m.bySKU[strings.ToUpper(sku)]
	if !exists {
		return nil, domain.ErrProductNotFound
	}
	return p, nil
}

func (m *mockProductRepository) GetAll(ctx context.Context, filter domain.ProductFilter) ([]domain.Product, int64, error) {
	list := make([]domain.Product, 0, len(m.products))
	for _, p := range m.products {
		list = append(list, *p)
	}
	return list, int64(len(list)), nil
}

func (m *mockProductRepository) Update(ctx context.Context, product *domain.Product) error {
	existing, exists := m.products[product.ID]
	if !exists {
		return domain.ErrProductNotFound
	}

	// SKU check
	newSKU := strings.ToUpper(product.SKU)
	if existing.SKU != newSKU {
		if _, exists := m.bySKU[newSKU]; exists {
			return domain.ErrDuplicateSKU
		}
		delete(m.bySKU, existing.SKU)
		m.bySKU[newSKU] = product
	}

	m.products[product.ID] = product
	return nil
}

func (m *mockProductRepository) Patch(ctx context.Context, id uint, updates map[string]interface{}) (*domain.Product, error) {
	p, exists := m.products[id]
	if !exists {
		return nil, domain.ErrProductNotFound
	}

	if name, ok := updates["name"].(string); ok {
		p.Name = name
	}
	if desc, ok := updates["description"].(string); ok {
		p.Description = desc
	}
	if sku, ok := updates["sku"].(string); ok {
		if strings.ToUpper(sku) != p.SKU {
			if _, exists := m.bySKU[strings.ToUpper(sku)]; exists {
				return nil, domain.ErrDuplicateSKU
			}
			delete(m.bySKU, p.SKU)
			p.SKU = strings.ToUpper(sku)
			m.bySKU[p.SKU] = p
		}
	}
	if price, ok := updates["price"].(float64); ok {
		p.Price = price
	}
	if stock, ok := updates["stock"].(int); ok {
		p.Stock = stock
	}
	if cat, ok := updates["category"].(string); ok {
		p.Category = cat
	}
	if status, ok := updates["status"].(string); ok {
		p.Status = status
	}

	return p, nil
}

func (m *mockProductRepository) Delete(ctx context.Context, id uint) error {
	p, exists := m.products[id]
	if !exists {
		return domain.ErrProductNotFound
	}
	delete(m.products, id)
	delete(m.bySKU, p.SKU)
	return nil
}
