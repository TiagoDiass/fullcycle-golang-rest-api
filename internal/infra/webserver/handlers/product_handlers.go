package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/dto"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/infra/database"
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
}
