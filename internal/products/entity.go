package products

import "gorm.io/gorm"

type ProductEntity struct {
	gorm.Model

	Name  string
	Price float64
	Stock int
}
