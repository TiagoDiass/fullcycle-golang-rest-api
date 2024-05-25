package database

import (
	"testing"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/stretchr/testify/require"
)

// func TestXxx(t *testing.T) {
// }

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
