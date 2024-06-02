package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// CreateProduct godoc
// @Summary      Creates a product
// @Description  Creates a product with name and price.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request		body			dto.CreateProductInput	true	"product params"
// @Success      201  			{object}  entity.Product
// @Failure      400  			{object}  Error
// @Failure      500  			{object}  Error
// @Router       /products 	[post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, req *http.Request) {
	var createProductDTO dto.CreateProductInput
	err := json.NewDecoder(req.Body).Decode(&createProductDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	product, err := entity.NewProduct(createProductDTO.Name, createProductDTO.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProductDB.Create(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
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

// ListProducts godoc
// @Summary      List products
// @Description  List products with pagination or with no pagination if page and limit are zero.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page		query			string	false		"page number"
// @Param        limit	query			string	false		"limit"
// @Success      200  	{array}		entity.Product
// @Failure      400  	{object}  Error
// @Failure      500  	{object}  Error
// @Router       /products 	[get]
// @Security ApiKeyAuth
func (h *ProductHandler) ListProducts(w http.ResponseWriter, req *http.Request) {
	page := req.URL.Query().Get("page")
	limit := req.URL.Query().Get("limit")
	sort := req.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
	return
}
