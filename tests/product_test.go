package tests

import (
	"testing"

	"github.com/DiegoHachig/ecommerce-management-system-go/internal/products"
)

func TestCreateProduct(t *testing.T) {

	product, err := products.NewProduct(
		1,
		"Laptop Lenovo",
		850,
		10,
	)

	if err != nil {
		t.Errorf("No se pudo crear el producto: %v", err)
	}

	if product.GetID() != 1 {
		t.Errorf("ID incorrecto")
	}

	if product.GetName() != "Laptop Lenovo" {
		t.Errorf("Nombre incorrecto")
	}
}
