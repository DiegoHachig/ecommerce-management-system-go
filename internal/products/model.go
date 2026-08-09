package products

import "errors"

// Product representa un producto del sistema.
// para la venta dentro del sistema de e-commerce.
type Product struct {
	id    int
	name  string
	price float64
	stock int
}

// Constructor
// NewProduct crea una nueva instancia de Product
// validando los datos recibidos antes de retornarla.
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
// GetID devuelve el identificador del producto.
func (p Product) GetID() int {
	return p.id
}

// GetName devuelve el nombre del producto.
func (p Product) GetName() string {
	return p.name
}

// GetPrice devuelve el precio del producto.
func (p Product) GetPrice() float64 {
	return p.price
}

// GetStock devuelve la cantidad disponible en inventario.
func (p Product) GetStock() int {
	return p.stock
}
