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

// func (r *ProductMap) FindByColorAndYear(vehicle domain.ProductAttributes) (v map[int]domain.Product, err error) {
// 	v = make(map[int]domain.Product)

// 	for key, value := range r.db {
// 		if vehicle.Color == value.Color && vehicle.FabricationYear == value.FabricationYear {
// 			v[key] = value
// 		}
// 	}

// 	if len(v) == 0 {
// 		return v, err
// 	}

// 	return v, nil
// }

// func (r *ProductMap) FindByBrandAndYearInterval(req internal.BrandYearRangeSearchType) (v map[int]domain.Product, err error) {
// 	v = make(map[int]domain.Product)

// 	for key, value := range r.db {
// 		if value.Brand == req.Brand {
// 			if value.FabricationYear >= req.StartYear && value.FabricationYear <= req.EndYear {
// 				v[key] = value
// 			}
// 		}
// 	}

// 	if len(v) == 0 {
// 		return v, err
// 	}

// 	return v, nil
// }

// func (r *ProductMap) GetAverageSpeedByBrand(b string) (v float64, err error) {
// 	var brandList []domain.Product
// 	for _, i := range r.db {
// 		if b == i.Brand {
// 			brandList = append(brandList, i)
// 		}
// 	}

// 	if len(brandList) == 0 {
// 		return 0, err
// 	}

// 	var sumSpeed float64

// 	for _, bl := range brandList {
// 		sumSpeed += bl.MaxSpeed
// 	}

// 	return sumSpeed / float64(len(brandList)), nil
// }

// func (r *ProductMap) Create(v domain.ProductAttributes) (err error) {
// 	vehicleList := make(map[int]domain.Product)

// 	maxKey := 0
// 	for _, v := range r.db {
// 		if v.Id > maxKey {
// 			maxKey = v.Id
// 		}
// 	}

// 	newID := maxKey + 1

// 	if _, exists := r.db[newID]; exists {
// 		return err
// 	}

// 	vehicleList[newID] = domain.Product{
// 		Id:                newID,
// 		VehicleAttributes: v,
// 	}

// 	return nil
// }

// func (r *ProductMap) CreateSome(vs []domain.ProductAttributes) (err error) {
// 	vehicleList := make(map[int]domain.Product)

// 	maxKey := 0
// 	for k := range r.db {
// 		if k > maxKey {
// 			maxKey = k
// 		}
// 	}

// 	for i, v := range vs {
// 		newID := maxKey + 1 + i

// 		if _, exists := r.db[newID]; exists {
// 			return err
// 		}

// 		vehicleList[i] = domain.Product{
// 			Id:                newID,
// 			VehicleAttributes: v,
// 		}
// 	}

// 	return nil
// }

// func (r *ProductMap) UpdateSpeed(v internal.UpdateSpeed) (err error) {
// 	var vehicle domain.Product

// 	for i := 0; i <= len(r.db); i++ {
// 		if r.db[i].Id == v.Id {
// 			vehicle = r.db[i]
// 			vehicle.MaxSpeed = v.Speed
// 			break
// 		}
// 	}

// 	if vehicle.Id == 0 {
// 		return err
// 	}

// 	return nil
// }

// func (r *ProductMap) GetByFuelType(t string) (v map[int]domain.Product, err error) {
// 	v = make(map[int]domain.Product)
// 	for key, i := range r.db {
// 		if i.FuelType == t {
// 			v[key] = i
// 		}
// 	}

// 	if len(v) == 0 {
// 		return v, err
// 	}

// 	fmt.Println("len", len(v))

// 	return v, nil
// }

// func (r *ProductMap) DeleteById(id int) (err error) {
// 	found := false
// 	db := r.db
// 	for key := range r.db {
// 		if key == id {
// 			delete(db, key)
// 			found = true
// 		}
// 	}

// 	if !found {
// 		return fmt.Errorf("404 Not Found: Veículo não encontrado.")
// 	}

// 	return nil
// }

// func (r *ProductMap) GetByTransmissionType(t string) (v map[int]domain.Product, err error) {
// 	v = make(map[int]domain.Product)

// 	for key, vs := range r.db {
// 		if vs.Transmission == t {
// 			v[key] = vs
// 		}
// 	}

// 	if len(v) == 0 {
// 		return v, fmt.Errorf("404 Not Found: Não foram encontrados veículos com esse tipo de transmissão.")
// 	}

// 	return v, err
// }

// func (r *ProductMap) UpdateFuelType(u internal.UpdateFuel) (err error) {
// 	found := false

// 	for _, vs := range r.db {
// 		if vs.Id == u.Id {
// 			found = true
// 			temp := vs
// 			temp.FuelType = u.FuelType
// 			fmt.Println("vehicle", temp)
// 		}
// 	}

// 	if !found {
// 		return fmt.Errorf("404 Not Found: Veículo não encontrado")
// 	}

// 	return nil
// }

// func (r *ProductMap) GetAverageCapacityByBrand(b string) (v float64, err error) {
// 	var sum int
// 	var list []domain.Product
// 	for _, i := range r.db {
// 		if i.Brand == b {
// 			sum += i.Capacity
// 			list = append(list, i)
// 		}
// 	}

// 	if len(list) == 0 {
// 		return 0, fmt.Errorf("404 Not Found: Não foram encontrados veículos dessa marca.")
// 	}

// 	v = float64(sum) / float64(len(list))

// 	return v, err
// }

// func (r *ProductMap) GetByDimensions(minLength, maxLength, minWidth, maxWidth float64) (v map[int]domain.Product, err error) {
// 	v = make(map[int]domain.Product)

// 	for key, i := range r.db {
// 		if i.Dimensions.Length >= minLength && i.Dimensions.Length <= maxLength {
// 			if i.Dimensions.Width >= minWidth && i.Dimensions.Width <= maxWidth {
// 				v[key] = i
// 			}
// 		}
// 	}

// 	if len(v) == 0 {
// 		return v, fmt.Errorf("404 Not Found: Não foram encontrados veículos com essas dimensões.")
// 	}

// 	return v, nil
// }

// func (r *ProductMap) GetByWeight(minW, maxW float64) (v map[int]domain.Product, err error) {
// 	v = make(map[int]domain.Product)
// 	for key, i := range r.db {
// 		if i.Weight >= minW && i.Weight <= maxW {
// 			v[key] = i
// 		}
// 	}

// 	if len(v) == 0 {
// 		return v, fmt.Errorf("404 Not Found: Não foram encontrados veículos nessa faixa de peso.")
// 	}

// 	return v, nil
// }
