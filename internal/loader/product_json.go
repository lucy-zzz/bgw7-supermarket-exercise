package loader

import (
	"app/internal/domain"
	"encoding/json"
	"os"
)

func NewProductJSONFile(path string) *ProductJSONFile {
	return &ProductJSONFile{
		path: path,
	}
}

type ProductJSONFile struct {
	path string
}

func (l *ProductJSONFile) Load() (p map[int]domain.Product, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var productJSON []domain.Product
	err = json.NewDecoder(file).Decode(&productJSON)
	if err != nil {
		return
	}

	p = make(map[int]domain.Product)
	for _, pr := range productJSON {
		p[pr.Id] = domain.Product{
			Id:          pr.Id,
			Name:        pr.Name,
			Quantity:    pr.Quantity,
			CodeValue:   pr.CodeValue,
			IsPublished: pr.IsPublished,
			Expiration:  pr.Expiration,
			Price:       pr.Price,
		}
	}

	return
}
