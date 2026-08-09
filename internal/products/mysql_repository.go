package products

import "github.com/DiegoHachig/ecommerce-management-system-go/database"

type MySQLRepository struct{}

// Save guarda un producto en MySQL.
func (r *MySQLRepository) Save(product Product) error {

	entity := ProductEntity{
		Name:  product.GetName(),
		Price: product.GetPrice(),
		Stock: product.GetStock(),
	}

	return database.DB.Create(&entity).Error
}

// GetAll obtiene todos los productos.
func (r *MySQLRepository) GetAll() ([]Product, error) {

	var entities []ProductEntity

	result := database.DB.Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	var products []Product

	for _, entity := range entities {

		product, _ := NewProduct(
			int(entity.ID),
			entity.Name,
			entity.Price,
			entity.Stock,
		)

		products = append(products, *product)
	}

	return products, nil
}
