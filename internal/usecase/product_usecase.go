package usecase

import (
	"fmt"

	"github.com/tumenbayev/go-products-api/internal/entity"
	"github.com/tumenbayev/go-products-api/internal/usecase/repository"
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
		return []entity.Product{}, fmt.Errorf("product usecase - get all: %w", err)
	}

	var filtered []entity.Product
	for _, p := range products {
		if filter.Category != "" && p.Category != filter.Category {
			continue
		}
		if filter.PriceLessThan > 0 && p.Price.Original > filter.PriceLessThan {
			continue
		}
		filtered = append(filtered, p)
	}

	for i, p := range filtered {
		discount := calculateDiscount(p)
		if discount > 0 {
			discountPercentage := fmt.Sprintf("%d%%", discount)
			filtered[i].Price.DiscountPercentage = &discountPercentage
			filtered[i].Price.Final = p.Price.Original - (p.Price.Original * discount / 100)
		} else {
			filtered[i].Price.Final = p.Price.Original
		}
	}

	if len(filtered) > 5 {
		filtered = filtered[:5]
	}

	return filtered, nil
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