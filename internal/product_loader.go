package internal

import "app/internal/domain"

type ProductLoader interface {
	Load() (v map[int]domain.Product, err error)
}
