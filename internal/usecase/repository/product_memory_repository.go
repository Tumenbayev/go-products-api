package repository

import (
	"sync"
	"github.com/tumenbayev/go-products-api/internal/entity"
)

type InMemoryProductRepository struct {
	products []entity.Product
	mu       sync.RWMutex
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: []entity.Product{
			{
				SKU:      "000001",
				Name:     "BV Lean leather ankle boots",
				Category: "boots",
				Price: entity.Price{
					Original: 89000,
					Final:    89000,
					Currency: "EUR",
				},
			},
			{
				SKU:      "000002",
				Name:     "BV Lean leather ankle boots",
				Category: "boots",
				Price: entity.Price{
					Original: 99000,
					Final:    99000,
					Currency: "EUR",
				},
			},
			{
				SKU:      "000003",
				Name:     "Ashlington leather ankle boots",
				Category: "boots",
				Price: entity.Price{
					Original: 71000,
					Final:    71000,
					Currency: "EUR",
				},
			},
			{
				SKU:      "000004",
				Name:     "Naima embellished suede sandals",
				Category: "sandals",
				Price: entity.Price{
					Original: 79500,
					Final:    79500,
					Currency: "EUR",
				},
			},
			{
				SKU:      "000005",
				Name:     "Nathane leather sneakers",
				Category: "sneakers",
				Price: entity.Price{
					Original: 59000,
					Final:    59000,
					Currency: "EUR",
				},
			},
		},
	}
}

func (r *InMemoryProductRepository) GetAll() ([]entity.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.products, nil
}