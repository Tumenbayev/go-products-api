package filters

import "github.com/tumenbayev/go-products-api/internal/entity"

type CategoryFilter struct {
    Category string
}

func (f *CategoryFilter) Apply(products []entity.Product) []entity.Product {
    var filtered []entity.Product
    for _, p := range products {
        if p.Category == f.Category {
            filtered = append(filtered, p)
        }
    }
    return filtered
}

func (f *CategoryFilter) Priority() int {
	return 1
}
