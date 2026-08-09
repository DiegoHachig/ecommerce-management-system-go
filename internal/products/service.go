package products

import "fmt"

// ProductService contiene la lógica de negocio
// relacionada con la gestión de productos.
type ProductService struct {
	repository ProductRepository
}

// NewProductService crea una nueva instancia
// del servicio de productos.
func NewProductService(repository ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) RegisterProduct(product Product) error {

	if product.GetPrice() <= 0 {
		return fmt.Errorf("precio inválido")
	}

	return s.repository.Save(product)
}

// MemoryRepository simula una base de datos
// almacenando productos temporalmente en memoria.
type MemoryRepository struct {
	products []Product
}

// Save almacena un producto en memoria.
func (r *MemoryRepository) Save(product Product) error {

	r.products = append(r.products, product)

	return nil
}

// GetAll devuelve todos los productos registrados.
func (r *MemoryRepository) GetAll() ([]Product, error) {
	return r.products, nil
}
