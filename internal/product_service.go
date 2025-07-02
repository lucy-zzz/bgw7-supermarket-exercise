package internal

import "app/internal/domain"

type ProductService interface {
	FindAll() (p map[int]domain.Product, err error)
	Create(p domain.Product) (err error)
	GetById(id int) (p domain.Product, err error)
	FindProducts(price float64) (p map[int]domain.Product, err error)
	UpdateById(id int, p domain.Product) (domain.Product, error)
	UpdateAttributesById(id int, p domain.Product) (domain.Product, error)
	DeleteById(id int) (err error)
}
