package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/lucy-zzz/bgw7-supermarket-exercise/internal/domain"
	"github.com/lucy-zzz/bgw7-supermarket-exercise/internal/dto"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("docs/db/products.json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to read products file"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRequestProducts
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	var products []domain.Product
	file, err := os.Open("docs/db/products.json")

	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	request := req.ToDomain()
	request.Id = len(products) + 1
	products = append(products, request)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	products, err := os.Open("docs/db/products.json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to read products file"}`))
		return
	}

	var decodedProducts []domain.Product

	json.NewDecoder(products).Decode(&decodedProducts)

	fmt.Println(decodedProducts)

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Product not found"}`))
		return
	}

	var product *domain.Product
	for _, d := range decodedProducts {
		if d.Id == id {
			product = &d
			break
		}
	}

	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Product not found"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	priceGtStr := r.URL.Query().Get("priceGt")
	if priceGtStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Missing parameter priceGt"}`))
		return
	}

	priceGt, err := strconv.ParseFloat(priceGtStr, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid value for priceGt"}`))
		return
	}

	file, err := os.Open("docs/db/products.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to read products file"}`))
		return
	}
	defer file.Close()

	var produtos []domain.Product
	json.NewDecoder(file).Decode(&produtos)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(`{"error": "Failed to parse products data"}`))
	// 	return
	// }

	var result []domain.Product
	for _, p := range produtos {
		if p.Price > priceGt {
			result = append(result, p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
