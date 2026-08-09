package products

import "errors"

// Product representa un producto del sistema.
type Product struct {
	id    int
	name  string
	price float64
	stock int
}

// Constructor
func NewProduct(id int, name string, price float64, stock int) (*Product, error) {

	if name == "" {
		return nil, errors.New("el nombre es obligatorio")
	}

	if price <= 0 {
		return nil, errors.New("el precio debe ser mayor a cero")
	}

	return &Product{
		id:    id,
		name:  name,
		price: price,
		stock: stock,
	}, nil
}

// Getters
func (p Product) GetID() int {
	return p.id
}

func (p Product) GetName() string {
	return p.name
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p Product) GetStock() int {
	return p.stock
}
