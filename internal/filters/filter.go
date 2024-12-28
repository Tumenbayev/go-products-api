package filters

import "github.com/tumenbayev/go-products-api/internal/entity"

type Filter interface {
    Apply(products []entity.Product) []entity.Product
    Priority() int
}