package handler

import (
	"app/internal"
	"app/internal/domain"
	"app/internal/dto"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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
		idStr := chi.URLParam(r, "id_product")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, "Malformed or incomplete product data.")
			return
		}

		data, err := h.sv.GetById(id)

		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
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

func (h *ProductDefault) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id_product")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		var input dto.CreateRequestProducts
		json.NewDecoder(r.Body).Decode(&input)

		prd := input.ToDomain()
		prd.Id = id

		data, err := h.sv.UpdateById(id, prd)

		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *ProductDefault) UpdateProductAttributes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id_product")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		var input dto.CreateRequestProducts
		json.NewDecoder(r.Body).Decode(&input)

		prd := input.ToDomain()
		prd.Id = id

		data, err := h.sv.UpdateById(id, prd)

		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *ProductDefault) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id_product")
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
