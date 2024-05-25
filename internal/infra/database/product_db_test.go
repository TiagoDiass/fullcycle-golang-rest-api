package database

import (
	"testing"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	db := getDatabaseConnection(t)

	product, _ := entity.NewProduct("Macbook M1", 7500)
	productDB := NewProductDB(db)

	err := productDB.Create(product)
	require.Nil(t, err)

	var createdProduct entity.Product

	err = db.First(&createdProduct, "id = ?", product.ID).Error
	require.Nil(t, err)
	require.Equal(t, createdProduct.ID, product.ID)
	require.Equal(t, createdProduct.Name, product.Name)
	require.Equal(t, createdProduct.Price, product.Price)
}

func TestFindProductById(t *testing.T) {
	db := getDatabaseConnection(t)

	createdProduct, _ := entity.NewProduct("Macbook M1", 7500)
	productDB := NewProductDB(db)

	err := productDB.Create(createdProduct)
	require.Nil(t, err)

	foundProduct, err := productDB.FindById(createdProduct.ID.String())

	require.Nil(t, err)
	require.NotNil(t, foundProduct)
	require.Equal(t, foundProduct.ID, createdProduct.ID)
	require.Equal(t, foundProduct.Name, createdProduct.Name)
}

func TestUpdateProduct(t *testing.T) {
	db := getDatabaseConnection(t)

	product, _ := entity.NewProduct("Macbook M1", 7500)
	productDB := NewProductDB(db)

	err := productDB.Create(product)
	require.Nil(t, err)

	product.Name = "iPhone 15"
	product.Price = 2500

	err = productDB.Update(product)
	require.Nil(t, err)

	updatedProduct, err := productDB.FindById(product.ID.String())
	require.Nil(t, err)

	require.Equal(t, updatedProduct.ID, product.ID)
	require.Equal(t, updatedProduct.Name, "iPhone 15")
	require.Equal(t, updatedProduct.Price, 2500)
}

func TestDeleteProduct(t *testing.T) {
	db := getDatabaseConnection(t)

	product, _ := entity.NewProduct("Macbook M1", 7500)
	productDB := NewProductDB(db)

	err := productDB.Create(product)
	require.Nil(t, err)

	err = productDB.Delete(product.ID.String())
	require.Nil(t, err)

	product, err = productDB.FindById(product.ID.String())
	require.Nil(t, product)
	require.NotNil(t, err)
}

func createAndSaveFiveProducts(t *testing.T, productDB *ProductDB) {
	partialProducts := []struct {
		name  string
		price int
	}{
		{name: "product 1", price: 1000},
		{name: "product 2", price: 2000},
		{name: "product 3", price: 3000},
		{name: "product 4", price: 4000},
		{name: "product 5", price: 5000},
	}

	for _, partialProduct := range partialProducts {
		product, err := entity.NewProduct(partialProduct.name, partialProduct.price)

		if err != nil {
			t.Fatalf("failed to create product: %v", err)
		}

		err = productDB.Create(product)

		if err != nil {
			t.Fatalf("failed to save product: %v", err)
		}
	}
}

func TestFindAllProductsWithPagination(t *testing.T) {
	db := getDatabaseConnection(t)
	productDB := NewProductDB(db)

	createAndSaveFiveProducts(t, productDB)

	result, err := productDB.FindAll(1, 2, "asc")
	require.Nil(t, err)
	require.Len(t, result, 2)

	require.Equal(t, result[0].Name, "product 1")
	require.Equal(t, result[1].Name, "product 2")

	result, err = productDB.FindAll(2, 2, "asc")

	require.Nil(t, err)
	require.Len(t, result, 2)

	require.Equal(t, result[0].Name, "product 3")
	require.Equal(t, result[1].Name, "product 4")

	result, err = productDB.FindAll(3, 2, "asc")

	require.Nil(t, err)
	require.Len(t, result, 1)
	require.Equal(t, result[0].Name, "product 5")
}

func TestFindAllProductsWithNoPagination(t *testing.T) {
	db := getDatabaseConnection(t)
	productDB := NewProductDB(db)

	createAndSaveFiveProducts(t, productDB)

	result, err := productDB.FindAll(0, 0, "asc")
	require.Nil(t, err)
	require.Len(t, result, 5)
	require.Equal(t, result[0].Name, "product 1")
	require.Equal(t, result[1].Name, "product 2")
	require.Equal(t, result[2].Name, "product 3")
	require.Equal(t, result[3].Name, "product 4")
	require.Equal(t, result[4].Name, "product 5")
}
