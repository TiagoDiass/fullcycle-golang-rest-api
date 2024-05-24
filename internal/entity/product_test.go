package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Macbook M1", 7500)

	require.Nil(t, err)
	require.NotNil(t, product)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, "Macbook M1")
	require.Equal(t, product.Price, 7500)
	require.NotEmpty(t, product.CreatedAt)
}

func TestProductWithNameRequiredError(t *testing.T) {
	product, err := NewProduct("", 7500)

	require.Nil(t, product)
	require.Equal(t, err, ErrNameIsRequired)
}

func TestProductWithPriceRequiredError(t *testing.T) {
	product, err := NewProduct("Macbook", 0)

	require.Nil(t, product)
	require.Equal(t, err, ErrPriceIsRequired)
}
