package products

// ProductRepository define las operaciones
// que debe implementar cualquier repositorio.
type ProductRepository interface {
	Save(product Product) error
	GetAll() ([]Product, error)
}
