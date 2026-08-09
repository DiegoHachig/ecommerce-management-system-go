package products

import "fmt"

// ShowProducts muestra los productos registrados.
func ShowProducts(products []Product) {

	for _, p := range products {

		fmt.Println(
			p.GetID(),
			p.GetName(),
			p.GetPrice(),
			p.GetStock(),
		)
	}
}
