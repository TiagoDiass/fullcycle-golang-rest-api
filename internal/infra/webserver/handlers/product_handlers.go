package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/dto"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/infra/database"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB database.IProductDB
}

func NewProductHandler(productDB database.IProductDB) *ProductHandler {
	return &ProductHandler{
		ProductDB: productDB,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {
	var createProductDTO dto.CreateProductInput
	err := json.NewDecoder(req.Body).Decode(&createProductDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entity.NewProduct(createProductDTO.Name, createProductDTO.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	return
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
	return
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updateProductDTO dto.UpdateProductInput
	err := json.NewDecoder(req.Body).Decode(&updateProductDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	product.Name = updateProductDTO.Name
	product.Price = updateProductDTO.Price

	err = h.ProductDB.Update(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
	return
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
