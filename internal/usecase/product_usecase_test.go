package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tumenbayev/go-products-api/internal/entity"
)

type mockProductRepository struct{}

func (m *mockProductRepository) GetAll() ([]entity.Product, error) {
	return []entity.Product{
		{SKU: "000001", Name: "BV Lean leather ankle boots", Category: "boots", Price: entity.Price{Original: 89000, Currency: "EUR"}},
		{SKU: "000002", Name: "BV Lean leather ankle boots", Category: "boots", Price: entity.Price{Original: 99000, Currency: "EUR"}},
		{SKU: "000003", Name: "Ashlington leather ankle boots", Category: "boots", Price: entity.Price{Original: 71000, Currency: "EUR"}},
		{SKU: "000004", Name: "Naima embellished suede sandals", Category: "sandals", Price: entity.Price{Original: 79500, Currency: "EUR"}},
		{SKU: "000005", Name: "Nathane leather sneakers", Category: "sneakers", Price: entity.Price{Original: 59000, Currency: "EUR"}},
	}, nil
}

func TestGetProducts(t *testing.T) {
	repo := &mockProductRepository{}
	useCase := NewProductUseCase(repo)

	t.Run("Filter by category", func(t *testing.T) {
		filter := ProductFilter{Category: "boots"}
		products, err := useCase.GetProducts(filter)
		assert.NoError(t, err)
		assert.Len(t, products, 3)
		assert.Equal(t, "000001", products[0].SKU)
		assert.Equal(t, "000002", products[1].SKU)
		assert.Equal(t, "000003", products[2].SKU)
	})

	t.Run("Filter by price less than", func(t *testing.T) {
		filter := ProductFilter{PriceLessThan: 80000}
		products, err := useCase.GetProducts(filter)
		assert.NoError(t, err)
		assert.Len(t, products, 3)
		assert.Equal(t, "000003", products[0].SKU)
		assert.Equal(t, "000004", products[1].SKU)
		assert.Equal(t, "000005", products[2].SKU)
	})

	t.Run("Apply discounts", func(t *testing.T) {
		filter := ProductFilter{}
		products, err := useCase.GetProducts(filter)
		assert.NoError(t, err)
		assert.Len(t, products, 5)

		assert.Equal(t, 62300, products[0].Price.Final) // 30% discount
		assert.Equal(t, "30%", *products[0].Price.DiscountPercentage)

		assert.Equal(t, 69300, products[1].Price.Final) // multiple discounts applied
		assert.Equal(t, "30%", *products[1].Price.DiscountPercentage)

		assert.Equal(t, 79500, products[3].Price.Final) // No discount
		assert.Nil(t, products[3].Price.DiscountPercentage)
	})

	t.Run("Limit to 5 products", func(t *testing.T) {
		filter := ProductFilter{}
		products, err := useCase.GetProducts(filter)
		assert.NoError(t, err)
		assert.Len(t, products, 5)
	})
}