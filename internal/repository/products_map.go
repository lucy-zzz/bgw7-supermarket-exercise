package repository

import (
	"github.com/lucy-zzz/bgw7-supermarket-exercise/internal/domain"
)

func NewProductsMap(db map[int]domain.Product) *ProductsMap {
	defaultDb := make(map[int]domain.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductsMap{db: defaultDb}
}

type ProductsMap struct {
	db map[int]domain.Product
}
