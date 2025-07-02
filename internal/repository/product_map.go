package repository

import (
	"app/internal/domain"
	"errors"
)

func NewProductMap(db map[int]domain.Product) *ProductMap {
	defaultDb := make(map[int]domain.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductMap{db: defaultDb}
}

type ProductMap struct {
	db map[int]domain.Product
}

func (m *ProductMap) FindAll() (p map[int]domain.Product, err error) {
	p = make(map[int]domain.Product)

	for key, value := range m.db {
		p[key] = value
	}

	return
}

func (m *ProductMap) Create(p domain.Product) (err error) {
	var new domain.Product

	id := len(m.db) + 1

	new.Id = id
	new.CodeValue = p.CodeValue
	new.Name = p.Name
	new.Expiration = p.Expiration
	new.IsPublished = p.IsPublished
	new.Price = p.Price
	new.Quantity = p.Quantity

	if m.db[id].Id != 0 {
		return errors.New("ID already exists.")
	}

	m.db[id] = new

	return nil
}

func (m *ProductMap) GetById(id int) (p domain.Product, err error) {
	var found domain.Product
	for _, value := range m.db {
		if value.Id == id {
			found = value
		}
	}

	if found.Id == 0 {
		return domain.Product{}, errors.New("ID not found.")
	}

	return found, nil
}

func (m *ProductMap) FindProducts(price float64) (p map[int]domain.Product, err error) {
	p = make(map[int]domain.Product)

	for key, pr := range m.db {
		if pr.Price > price {
			p[key] = pr
		}
	}

	if len(p) == 0 {
		return p, errors.New("Not found.")
	}

	return p, nil
}

func (m *ProductMap) DeleteById(id int) (err error) {
	list := make(map[int]domain.Product)
	var found bool
	for key, value := range m.db {
		if value.Id != id {
			list[key] = value
		} else {
			found = true
		}
	}

	m.db = map[int]domain.Product{}
	m.db = list

	if !found {
		return errors.New("ID not found.")
	}

	return err
}

func (m *ProductMap) UpdateById(id int, p domain.Product) (r domain.Product, err error) {
	var found bool
	for key, value := range m.db {
		if value.Id == id {
			m.db[key] = p
			r = m.db[key]
			found = true
		}
	}

	if !found {
		return r, errors.New("ID not found.")
	}

	return r, err
}

func (m *ProductMap) UpdateAttributesById(id int, p domain.Product) (r domain.Product, err error) {
	var found bool
	for key, value := range m.db {
		if value.Id == id {
			product := m.db[key]

			if p.CodeValue != "" {
				product.CodeValue = p.CodeValue
			}

			if p.Quantity != 0 {
				product.Quantity = p.Quantity
			}

			if p.Price != 0 {
				product.Price = p.Price
			}

			if p.IsPublished != product.IsPublished {
				product.IsPublished = p.IsPublished
			}

			if p.Expiration != "" {
				product.Expiration = p.Expiration
			}

			if p.Name != "" {
				product.Name = p.Name
			}

			m.db[key] = product
			found = true
		}
	}

	if !found {
		return r, errors.New("ID not found.")
	}

	return r, err
}
