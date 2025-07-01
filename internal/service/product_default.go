package service

import (
	"app/internal"
	"app/internal/domain"
)

func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{rp: rp}
}

type ProductDefault struct {
	rp internal.ProductRepository
}

func (s *ProductDefault) FindAll() (v map[int]domain.Product, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *ProductDefault) Create(new domain.Product) (err error) {
	err = s.rp.Create(new)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProductDefault) GetById(id int) (p domain.Product, err error) {
	return s.rp.GetById(id)
}

func (s *ProductDefault) FindProducts(price float64) (p map[int]domain.Product, err error) {
	return s.rp.FindProducts(price)
}

func (s *ProductDefault) DeleteById(id int) error {
	return s.rp.DeleteById(id)
}

// func (s *ProductDefault) FindByColorAndYear(vehicle domain.Product) (v map[int]domain.Product, err error) {
// 	v, err = s.rp.FindByColorAndYear(vehicle)

// 	if err != nil {
// 		return v, err
// 	}

// 	return v, nil
// }

// func (s *ProductDefault) FindByBrandAndYearInterval(r internal.BrandYearRangeSearchType) (v map[int]domain.Product, err error) {
// 	v, err = s.rp.FindByBrandAndYearInterval(r)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return v, nil
// }

// func (s *ProductDefault) GetAverageSpeedByBrand(b string) (v float64, err error) {
// 	v, err = s.rp.GetAverageSpeedByBrand(b)

// 	if err != nil {
// 		return 0, err
// 	}

// 	return v, nil
// }

// func (s *ProductDefault) CreateSome(vs []domain.Product) (err error) {
// 	err = s.rp.CreateSome(vs)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *ProductDefault) UpdateSpeed(v internal.UpdateSpeed) (err error) {
// 	err = s.rp.UpdateSpeed(v)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *ProductDefault) GetByFuelType(t string) (v map[int]domain.Product, err error) {
// 	v, err = s.rp.GetByFuelType(t)

// 	if err != nil {
// 		return v, err
// 	}

// 	return v, nil
// }

// func (s *ProductDefault) GetByTransmissionType(t string) (v map[int]domain.Product, err error) {
// 	v, err = s.rp.GetByTransmissionType(t)

// 	return v, err
// }

// func (s *ProductDefault) UpdateFuelType(u internal.UpdateFuel) (err error) {
// 	err = s.rp.UpdateFuelType(u)
// 	return err
// }

// func (s *ProductDefault) GetAverageCapacityByBrand(b string) (v float64, err error) {
// 	v, err = s.rp.GetAverageCapacityByBrand(b)

// 	return v, err
// }

// func (s *ProductDefault) GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]domain.Product, err error) {
// 	v, err = s.rp.GetByDimensions(minLength, maxLength, minWidth, maxWidth)

// 	return v, err
// }

// func (s *ProductDefault) GetByWeight(minW, maxW float64) (v map[int]domain.Product, err error) {
// 	v, err = s.rp.GetByWeight(minW, maxW)

// 	return v, err
// }
