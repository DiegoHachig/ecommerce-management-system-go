package products

import "github.com/DiegoHachig/ecommerce-management-system-go/database"

// MySQLRepository implementa el acceso a datos
// utilizando MySQL y GORM para la persistencia
// de productos.
type MySQLRepository struct{}

// Save almacena un producto en la base de datos
// convirtiendo el modelo de negocio en una entidad GORM.

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
