package usecase

import (
	"fmt"
	"sort"

	"github.com/tumenbayev/go-products-api/internal/entity"
	"github.com/tumenbayev/go-products-api/internal/usecase/repository"
	"github.com/tumenbayev/go-products-api/internal/filters"
)

type ProductUseCase interface {
	GetProducts(filter ProductFilter) ([]entity.Product, error)
}

type ProductFilter struct {
	Category      string
	PriceLessThan int
}

type productUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: repo,
	}
}

func (uc *productUseCase) GetProducts(filter ProductFilter) ([]entity.Product, error) {
	products, err := uc.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("product usecase - get all: %w", err)
	}

	var activeFilters []filters.Filter
	if filter.Category != "" {
		activeFilters = append(activeFilters, &filters.CategoryFilter{Category: filter.Category})
	}
	if filter.PriceLessThan > 0 {
		activeFilters = append(activeFilters, &filters.PriceFilter{PriceLessThan: filter.PriceLessThan})
	}

	sort.Slice(activeFilters, func(i, j int) bool {
		return activeFilters[i].Priority() < activeFilters[j].Priority()
	})

	for _, f := range activeFilters {
		products = f.Apply(products)
	}

	for i := range products {
		discount := calculateDiscount(products[i])
		if discount > 0 {
			discountPercentage := fmt.Sprintf("%d%%", discount)
			products[i].Price.DiscountPercentage = &discountPercentage
			products[i].Price.Final = products[i].Price.Original - (products[i].Price.Original * discount / 100)
		} else {
			products[i].Price.Final = products[i].Price.Original
		}
	}

	if len(products) > 5 {
		products = products[:5]
	}

	return products, nil
}

func calculateDiscount(p entity.Product) int {
	discounts := []int{}

	if p.Category == "boots" {
		discounts = append(discounts, 30)
	}

	if p.SKU == "000003" {
		discounts = append(discounts, 15)
	}

	if len(discounts) > 0 {
		return max(discounts)
	}
	return 0
}

func max(nums []int) int {
	maxVal := nums[0]
	for _, n := range nums[1:] {
		if n > maxVal {
			maxVal = n
		}
	}
	return maxVal
}