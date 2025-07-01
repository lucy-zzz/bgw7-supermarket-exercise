package handler

import (
	"app/internal"
	"app/internal/domain"
	"app/internal/dto"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
)

func NewProductDefault(sv internal.ProductService) *ProductDefault {
	return &ProductDefault{sv: sv}
}

type ProductDefault struct {
	sv internal.ProductService
}

func (h *ProductDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := make(map[int]domain.Product)
		for key, value := range v {
			data[key] = domain.Product{
				Id:          value.Id,
				Name:        value.Name,
				Quantity:    value.Quantity,
				CodeValue:   value.CodeValue,
				IsPublished: value.IsPublished,
				Expiration:  value.Expiration,
				Price:       value.Price,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *ProductDefault) CreateProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody dto.CreateRequestProducts

		if err := request.JSON(r, &requestBody); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		validate := validator.New()
		if err := validate.Struct(requestBody); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, "validation error")
			return
		}

		err := h.sv.Create(requestBody.ToDomain())

		if err != nil {
			response.Error(w, http.StatusConflict, err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "created",
		})
	}
}

func (h *ProductDefault) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
			response.JSON(w, http.StatusBadRequest, 400)
			return
		}

		data, err := h.sv.GetById(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *ProductDefault) SearchProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGtStr := r.URL.Query().Get("priceGt")

		if priceGtStr == "" {
			response.Error(w, http.StatusBadRequest, "Missing parameter priceGt")
			return
		}

		priceGt, err := strconv.ParseFloat(priceGtStr, 64)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		data, err := h.sv.FindProducts(priceGt)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *ProductDefault) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		err = h.sv.DeleteById(id)

		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "204 No Content: Veículo removido com sucesso.",
		})
	}
}

// func CreateProducts(w http.ResponseWriter, r *http.Request) {
// 	var req dto.CreateRequestProducts
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		response.Error(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	var products []domain.Product
// 	file, err := os.Open("docs/db/products.json")

// 	if err != nil {
// 		response.Error(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// func GetProducts(w http.ResponseWriter, r *http.Request) {
// 	data, err := os.ReadFile("docs/db/products.json")

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"error": "Failed to read products file"}`))
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(data)
// }

// 	err = json.NewDecoder(file).Decode(&products)
// 	if err != nil {
// 		response.Error(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	request := req.ToDomain()
// 	request.Id = len(products) + 1
// 	products = append(products, request)
// }

// func GetProductById(w http.ResponseWriter, r *http.Request) {
// 	products, err := os.Open("docs/db/products.json")

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"error": "Failed to read products file"}`))
// 		return
// 	}

// 	var decodedProducts []domain.Product

// 	json.NewDecoder(products).Decode(&decodedProducts)

// 	fmt.Println(decodedProducts)

// 	idStr := chi.URLParam(r, "id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`{"error": "Product not found"}`))
// 		return
// 	}

// 	var product *domain.Product
// 	for _, d := range decodedProducts {
// 		if d.Id == id {
// 			product = &d
// 			break
// 		}
// 	}

// 	if product == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`{"error": "Product not found"}`))
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(product)
// }

// func SearchProducts(w http.ResponseWriter, r *http.Request) {
// 	priceGtStr := r.URL.Query().Get("priceGt")
// 	if priceGtStr == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(`{"error": "Missing parameter priceGt"}`))
// 		return
// 	}

// 	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(`{"error": "Invalid value for priceGt"}`))
// 		return
// 	}

// 	file, err := os.Open("docs/db/products.json")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"error": "Failed to read products file"}`))
// 		return
// 	}
// 	defer file.Close()

// 	var produtos []domain.Product
// 	json.NewDecoder(file).Decode(&produtos)
// 	// if err != nil {
// 	// 	w.WriteHeader(http.StatusInternalServerError)
// 	// 	w.Write([]byte(`{"error": "Failed to parse products data"}`))
// 	// 	return
// 	// }

// 	var result []domain.Product
// 	for _, p := range produtos {
// 		if p.Price > priceGt {
// 			result = append(result, p)
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(result)
// }

///////

// func (h *ProductDefault) Create() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var input domain.Product

// 		err := json.NewDecoder(r.Body).Decode(&input)

// 		if err != nil {
// 			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		newProductAAtributes := domain.Product{
// 			Id:          input.Id,
// 			Name:        input.Name,
// 			Quantity:    input.Quantity,
// 			CodeValue:   input.CodeValue,
// 			IsPublished: input.IsPublished,
// 			Expiration:  input.Expiration,
// 			Price:       input.Price,
// 		}

// 		err = h.sv.Create(newProductAAtributes)
// 		if err != nil {
// 			w.Write([]byte(`{message: 409 Conflict: Identificador do veículo já existente.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		response.JSON(w, http.StatusCreated, map[string]any{
// 			"message": "success",
// 		})

// 	}
// }

// func (h *ProductDefault) GetByColorAndYear() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		color := r.URL.Query().Get("color")
// 		yearStr := r.URL.Query().Get("year")
// 		year, err := strconv.Atoi(yearStr)

// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(`{"message":"Parâmetro year inválido ou ausente."}`))
// 			return
// 		}

// 		input := domain.Product{
// 			Color:           color,
// 			FabricationYear: year,
// 		}

// 		vehicle := domain.Product{
// 			Color:           input.Color,
// 			FabricationYear: input.FabricationYear,
// 		}

// 		vehiclesList, err := h.sv.FindByColorAndYear(vehicle)

// 		if err != nil {
// 			w.Write([]byte(`{"message": "404 Not Found: Nenhum veículo encontrado com esses critérios." }`))
// 			response.JSON(w, http.StatusNotFound, 404)
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "success",
// 			"data":    vehiclesList,
// 		})
// 	}
// }

// func (h *ProductDefault) GetByBrandAndYearInterval() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		brand := chi.URLParam(r, "brand")
// 		startYearStr := chi.URLParam(r, "start_year")
// 		endYearStr := chi.URLParam(r, "end_year")

// 		startYear, err := strconv.Atoi(startYearStr)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(`{"message":"Parâmetro start_year inválido ou ausente."}`))
// 			return
// 		}

// 		endYear, err := strconv.Atoi(endYearStr)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(`{"message":"Parâmetro end_year inválido ou ausente."}`))
// 			return
// 		}

// 		req := internal.BrandYearRangeSearchType{
// 			Brand:     brand,
// 			StartYear: startYear,
// 			EndYear:   endYear,
// 		}

// 		vehiclesList, err := h.sv.FindByBrandAndYearInterval(req)

// 		if err != nil {
// 			w.Write([]byte(`{"message": "404 Not Found: Nenhum veículo encontrado com esses critérios." }`))
// 			response.JSON(w, http.StatusNotFound, 404)
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "success",
// 			"data":    vehiclesList,
// 		})
// 	}
// }

// func (h *ProductDefault) GetAverageSpeedByBrand() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		brand := chi.URLParam(r, "brand")

// 		averageSpeed, err := h.sv.GetAverageSpeedByBrand(brand)

// 		if err != nil {
// 			w.Write([]byte(`{message: 404 Not Found: Nenhum veículo encontrado dessa marca.}`))
// 			response.JSON(w, http.StatusNotFound, 404)
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "success",
// 			"data":    averageSpeed})

// 	}
// }

// func (h *ProductDefault) CreateSome() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var input []domain.Product
// 		err := json.NewDecoder(r.Body).Decode(&input)

// 		if err != nil {
// 			w.Write([]byte(`"message": "400 Bad Request: Dados de algum veículo malformados ou incompletos."`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		err = h.sv.CreateSome(input)

// 		if err != nil {
// 			w.Write([]byte(`"message": "409 Conflict: Algum veículo possui um identificador já existente."`))
// 			response.JSON(w, http.StatusConflict, 409)
// 			return
// 		}

// 		response.JSON(w, http.StatusCreated, 201)
// 	}
// }

// func (h *ProductDefault) UpdateSpeed() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var s struct {
// 			Speed float64 `json:"speed"`
// 		}

// 		err := json.NewDecoder(r.Body).Decode(&s)

// 		if err != nil {
// 			response.JSON(w, http.StatusBadRequest, `{"message": "400 Bad Request: Velocidade malformada ou fora de alcance."}`)
// 			return
// 		}

// 		idStr := chi.URLParam(r, "id")
// 		id, err := strconv.Atoi(idStr)

// 		if err != nil {
// 			w.Write([]byte(`400 Bad Request: Velocidade malformada ou fora de alcance.`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		var u internal.UpdateSpeed
// 		u = internal.UpdateSpeed{
// 			Id:    int(id),
// 			Speed: s.Speed,
// 		}

// 		err = h.sv.UpdateSpeed(u)

// 		if err != nil {
// 			w.Write([]byte(`404 Not Found: Veículo não encontrado.`))
// 			response.JSON(w, http.StatusNotFound, 404)
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, 200)
// 	}
// }

// func (h *ProductDefault) GetByFuelType() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fuelType := chi.URLParam(r, "type")

// 		data, err := h.sv.GetByFuelType(fuelType)

// 		if err != nil {
// 			w.Write([]byte(`404 Not Found: Não foram encontrados veículos com esse tipo de combustível.`))
// 			response.JSON(w, http.StatusNotFound, 404)
// 			return
// 		}

// 		if len(data) == 0 {
// 			w.WriteHeader(http.StatusNotFound)
// 			response.JSON(w, http.StatusNotFound, map[string]string{
// 				"message": "404 Not Found: Não foram encontrados veículos com esse tipo de combustível.",
// 			})
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "success",
// 			"data":    data,
// 		})
// 		return
// 	}
// }

// func (h *ProductDefault) DeleteById() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		idStr := chi.URLParam(r, "id")

// 		id, err := strconv.Atoi(idStr)

// 		if err != nil {
// 			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		err = h.sv.DeleteById(id)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusNotFound)
// 			return
// 		}

// 		response.JSON(w, http.StatusNoContent, map[string]any{
// 			"message": "204 No Content: Veículo removido com sucesso.",
// 		})
// 	}
// }

// func (h *ProductDefault) GetByTransmissionType() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		t := chi.URLParam(r, "type")

// 		data, err := h.sv.GetByTransmissionType(t)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusNotFound)
// 			w.Write([]byte(`{message: 404 Not Found: Não foram encontrados veículos com esse tipo de transmissão.}`))
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "success",
// 			"data":    data,
// 		})
// 	}
// }

// func (h *ProductDefault) UpdateFuel() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		idStr := chi.URLParam(r, "id")
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			http.Error(w, err.Error(), 400)
// 			w.Write([]byte(`400 Bad Request: Tipo de combustível malformado ou não suportado.`))
// 			return
// 		}

// 		var f struct {
// 			FuelType string `json:"fuel_type"`
// 		}

// 		err = json.NewDecoder(r.Body).Decode(&f)
// 		if err != nil {
// 			http.Error(w, err.Error(), 400)
// 			w.Write([]byte(`400 Bad Request: Tipo de combustível malformado ou não suportado.`))
// 			return
// 		}

// 		v := internal.UpdateFuel{
// 			Id:       id,
// 			FuelType: f.FuelType,
// 		}

// 		err = h.sv.UpdateFuelType(v)

// 		if err != nil {
// 			http.Error(w, err.Error(), 404)
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, 200)
// 	}
// }

// func (h *ProductDefault) GetAverageCapacityByBrand() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		brand := chi.URLParam(r, "brand")

// 		data, err := h.sv.GetAverageCapacityByBrand(brand)

// 		if err != nil {
// 			http.Error(w, err.Error(), 404)
// 			w.Write([]byte(`"message": "404 Not Found: Não foram encontrados veículos dessa marca."`))
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "success",
// 			"data":    data,
// 		})
// 	}
// }

// func (h *ProductDefault) GetByDimensions() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		length := r.URL.Query().Get("length")

// 		l := strings.Split(length, "-")

// 		minLength, err := strconv.ParseFloat(l[0], 64)

// 		if err != nil {
// 			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		maxLength, err := strconv.ParseFloat(l[1], 64)

// 		if err != nil {
// 			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		width := r.URL.Query().Get("width")
// 		wh := strings.Split(width, "-")

// 		minWidth, err := strconv.ParseFloat(wh[0], 64)

// 		if err != nil {
// 			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		maxWidth, err := strconv.ParseFloat(wh[1], 64)

// 		if err != nil {
// 			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		data, err := h.sv.GetByDimensions(minLength, maxLength, minWidth, maxWidth)

// 		if err != nil {
// 			w.Write([]byte(`{message: 404 Not Found: Não foram encontrados veículos com essas dimensões.}`))
// 			response.JSON(w, http.StatusNotFound, 404)
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "success",
// 			"data":    data,
// 		})

// 	}
// }

// func (h *ProductDefault) GetByWeight() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		min := r.URL.Query().Get("min")
// 		max := r.URL.Query().Get("max")

// 		wmin, err := strconv.ParseFloat(min, 64)

// 		if err != nil {
// 			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		wmax, err := strconv.ParseFloat(max, 64)

// 		if err != nil {
// 			w.Write([]byte(`{message: 400 Bad Request: Dados do veículo mal formatados ou incompletos.}`))
// 			response.JSON(w, http.StatusBadRequest, 400)
// 			return
// 		}

// 		data, err := h.sv.GetByWeight(wmin, wmax)

// 		if err != nil {
// 			w.Write([]byte(`{message: 404 Not Found: Não foram encontrados veículos nessa faixa de peso.}`))
// 			response.JSON(w, http.StatusNotFound, 404)
// 			return
// 		}

// 		response.JSON(w, http.StatusOK, map[string]any{
// 			"message": "success",
// 			"data":    data,
// 		})

// 	}
// }
