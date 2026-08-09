package products

import "fmt"

type ProductService struct {
	repository ProductRepository
}

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

type MemoryRepository struct {
	products []Product
}

func (r *MemoryRepository) Save(product Product) error {

	r.products = append(r.products, product)

	return nil
}

func (r *MemoryRepository) GetAll() ([]Product, error) {
	return r.products, nil
}
