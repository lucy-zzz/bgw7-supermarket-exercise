package handler_test

import (
	"app/internal/domain"
	"app/internal/dto"
	"app/internal/handler"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type mockProductService struct {
	FindAllFunc              func() (map[int]domain.Product, error)
	CreateFunc               func(p domain.Product) (err error)
	GetByIdFunc              func(id int) (p domain.Product, err error)
	FindProductsFunc         func(price float64) (p map[int]domain.Product, err error)
	UpdateByIdFunc           func(id int, pr domain.Product) (p domain.Product, e error)
	UpdateAttributesByIdFunc func(id int, pr domain.Product) (p domain.Product, e error)
	DeleteByIdFunc           func(id int) (err error)
}

func (m *mockProductService) Create(p domain.Product) error {
	return nil
}

func (m *mockProductService) GetById(id int) (domain.Product, error) {
	return domain.Product{}, nil
}

func (m *mockProductService) DeleteById(id int) error {
	return nil
}

func (m *mockProductService) UpdateById(id int, p domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}

func (m *mockProductService) FindProducts(price float64) (map[int]domain.Product, error) {
	products := map[int]domain.Product{
		1: {Id: 1, Name: "Product 1", Price: 100.0},
		2: {Id: 2, Name: "Product 2", Price: 200.0},
	}

	filteredProducts := make(map[int]domain.Product)
	for id, product := range products {
		if product.Price >= price {
			filteredProducts[id] = product
		}
	}

	return filteredProducts, nil
}

func (m *mockProductService) FindAll() (map[int]domain.Product, error) {
	return m.FindAllFunc()
}

func (m *mockProductService) UpdateAttributesById(id int, p domain.Product) (domain.Product, error) {
	if m.UpdateAttributesByIdFunc != nil {
		return m.UpdateAttributesByIdFunc(id, p)
	}
	return domain.Product{}, nil
}

func TestGetAll_Success(t *testing.T) {
	mockProducts := map[int]domain.Product{
		1: {Id: 1, Name: "Produto 1", Quantity: 3, CodeValue: "123", IsPublished: true, Expiration: "2025-01-01", Price: 10.0},
		2: {Id: 2, Name: "Produto 2", Quantity: 5, CodeValue: "456", IsPublished: false, Expiration: "2025-06-01", Price: 20.0},
	}

	mockSvc := &mockProductService{
		FindAllFunc: func() (map[int]domain.Product, error) {
			return mockProducts, nil
		},
	}
	pHandler := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	pHandler.GetAll().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAll_Failure(t *testing.T) {
	mockSvc := &mockProductService{
		FindAllFunc: func() (map[int]domain.Product, error) {
			return nil, errors.New("database failure")
		},
	}
	pHandler := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	pHandler.GetAll().ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateProducts_Success(t *testing.T) {
	mockService := &mockProductService{
		CreateFunc: func(p domain.Product) error {
			return nil
		},
	}

	handler := handler.NewProductDefault(mockService)

	requestPayload := dto.CreateRequestProducts{
		Name:        "Test Product",
		Quantity:    10,
		CodeValue:   "ABC123",
		IsPublished: true,
		Expiration:  "2025-01-01",
		Price:       15.5,
	}

	bodyBytes, err := json.Marshal(requestPayload)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateProducts().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var respBody map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, "created", respBody["message"])
}

func TestCreateProducts_BadRequest(t *testing.T) {
	mockService := &mockProductService{}

	h := handler.NewProductDefault(mockService)

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte("not json")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	h.CreateProducts().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestGetProductById_Success(t *testing.T) {
	mockProduct := domain.Product{
		Id:          1,
		Name:        "Produto Teste",
		Quantity:    10,
		CodeValue:   "ABC123",
		IsPublished: true,
		Expiration:  "2025-01-01",
		Price:       99.99,
	}

	mockSvc := &mockProductService{
		GetByIdFunc: func(id int) (domain.Product, error) {
			return mockProduct, nil
		},
	}

	h := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_product", strconv.Itoa(mockProduct.Id))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h.GetProductById().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductById_BadRequest(t *testing.T) {
	mockSvc := &mockProductService{}

	h := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/products/abc", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_product", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h.GetProductById().ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSearchProducts_Success(t *testing.T) {
	mockProducts := []domain.Product{
		{Id: 1, Name: "Produto A", Price: 10.5},
		{Id: 2, Name: "Produto B", Price: 20.0},
	}

	mockSvc := &mockProductService{
		FindProductsFunc: func(price float64) (map[int]domain.Product, error) {
			assert.Equal(t, 10.0, price)
			result := make(map[int]domain.Product)
			for i, p := range mockProducts {
				result[i+1] = p
			}
			return result, nil
		},
	}

	h := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/products/search?priceGt=10.0", nil)
	w := httptest.NewRecorder()

	h.SearchProducts().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp["message"])
}

func TestSearchProducts_MissingPriceGt(t *testing.T) {
	mockSvc := &mockProductService{}
	h := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/products/search", nil)
	w := httptest.NewRecorder()

	h.SearchProducts().ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSearchProducts_InvalidPriceGt(t *testing.T) {
	mockSvc := &mockProductService{}
	h := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/products/search?priceGt=abc", nil)
	w := httptest.NewRecorder()

	h.SearchProducts().ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateProduct_Success(t *testing.T) {
	mockUpdatedProduct := domain.Product{
		Id:    1,
		Name:  "Produto Atualizado",
		Price: 99.99,
	}

	mockSvc := &mockProductService{
		UpdateByIdFunc: func(id int, p domain.Product) (domain.Product, error) {
			assert.Equal(t, 1, id)
			assert.Equal(t, "Produto Atualizado", p.Name)
			return mockUpdatedProduct, nil
		},
	}

	h := handler.NewProductDefault(mockSvc)

	body := dto.CreateRequestProducts{
		Name:        "Produto Atualizado",
		Quantity:    5,
		CodeValue:   "X123",
		IsPublished: true,
		Expiration:  "2026-01-01",
		Price:       99.99,
	}

	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_product", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h.UpdateProduct().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp["message"])
}

func TestUpdateProduct_BadRequest(t *testing.T) {
	mockSvc := &mockProductService{}

	h := handler.NewProductDefault(mockSvc)

	body := dto.CreateRequestProducts{}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPut, "/products/abc", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_product", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h.UpdateProduct().ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateProductAttributes_Success(t *testing.T) {
	mockUpdatedProduct := domain.Product{
		Id:    2,
		Name:  "Produto Modificado",
		Price: 55.55,
	}

	mockSvc := &mockProductService{
		UpdateByIdFunc: func(id int, p domain.Product) (domain.Product, error) {
			assert.Equal(t, 2, id)
			assert.Equal(t, "Produto Modificado", p.Name)
			return mockUpdatedProduct, nil
		},
	}

	h := handler.NewProductDefault(mockSvc)

	body := dto.CreateRequestProducts{
		Name:        "Produto Modificado",
		Quantity:    3,
		CodeValue:   "Z987",
		IsPublished: false,
		Expiration:  "2024-12-31",
		Price:       55.55,
	}

	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPatch, "/products/2", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_product", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h.UpdateProductAttributes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "success", resp["message"])
}

func TestUpdateProductAttributes_BadRequest(t *testing.T) {
	mockSvc := &mockProductService{}

	h := handler.NewProductDefault(mockSvc)

	body := dto.CreateRequestProducts{}
	bodyBytes, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPatch, "/products/abc", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_product", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h.UpdateProductAttributes().ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteProduct_Success(t *testing.T) {
	mockSvc := &mockProductService{
		DeleteByIdFunc: func(id int) error {
			assert.Equal(t, 1, id)
			return nil
		},
	}

	h := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_product", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h.DeleteProduct().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteProduct_BadRequest(t *testing.T) {
	mockSvc := &mockProductService{}

	h := handler.NewProductDefault(mockSvc)

	req := httptest.NewRequest(http.MethodDelete, "/products/abc", nil)
	w := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_product", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h.DeleteProduct().ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
