package orders

import "fmt"

// OrderService contiene la lógica de negocio.
type OrderService struct {
	repository OrderRepository
}

// NewOrderService crea una nueva instancia del servicio.
func NewOrderService(repository OrderRepository) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

// CreateOrder registra un nuevo pedido.
func (s *OrderService) CreateOrder(order Order) error {

	if order.GetTotal() <= 0 {
		return fmt.Errorf("total invalido")
	}

	return s.repository.Save(order)
}

// MemoryRepository simula una base de datos en memoria.
type MemoryRepository struct {
	orders []Order
}

// Save almacena un pedido.
func (r *MemoryRepository) Save(order Order) error {

	r.orders = append(r.orders, order)

	return nil
}

// GetAll devuelve todos los pedidos registrados.
func (r *MemoryRepository) GetAll() ([]Order, error) {
	return r.orders, nil
}
