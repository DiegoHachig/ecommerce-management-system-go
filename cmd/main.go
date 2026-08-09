package main

import (
	"fmt"

	"github.com/DiegoHachig/ecommerce-management-system-go/internal/products"
)

func main() {

	repo := &products.MemoryRepository{}

	service := products.NewProductService(repo)

	product, err := products.NewProduct(
		1,
		"Laptop Lenovo",
		850,
		10,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = service.RegisterProduct(*product)

	if err != nil {
		fmt.Println(err)
		return
	}

	list, _ := repo.GetAll()

	products.ShowProducts(list)
}
