package dto

import (
	"app/internal/domain"
	"time"
)

type CreateRequestProducts struct {
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

func (c CreateRequestProducts) ToDomain() domain.Product {
	return domain.Product{
		Name:        c.Name,
		Quantity:    c.Quantity,
		CodeValue:   c.CodeValue,
		IsPublished: c.IsPublished,
		Expiration:  c.Expiration,
		Price:       c.Price,
	}
}
