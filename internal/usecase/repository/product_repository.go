package repository

import "github.com/tumenbayev/go-products-api/internal/entity"

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
}