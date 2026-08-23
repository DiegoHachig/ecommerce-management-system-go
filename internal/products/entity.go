package products

import "gorm.io/gorm"

// ProductEntity representa la estructura de la tabla
// de productos almacenada en MySQL mediante GORM.
type ProductEntity struct {
	gorm.Model

	Name  string
	Price float64
	Stock int
}
