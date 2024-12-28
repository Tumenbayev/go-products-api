package filters

import "github.com/tumenbayev/go-products-api/internal/entity"

type PriceFilter struct {
    PriceLessThan int
}

func (f *PriceFilter) Apply(products []entity.Product) []entity.Product {
    var filtered []entity.Product
    for _, p := range products {
        if p.Price.Original < f.PriceLessThan {
            filtered = append(filtered, p)
        }
    }
    return filtered
}

func (f *PriceFilter) Priority() int {
    return 2
}