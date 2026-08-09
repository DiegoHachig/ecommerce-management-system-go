package inventory

import "fmt"

// InventoryService contiene la lógica de negocio.
type InventoryService struct {
	repository InventoryRepository
}

// NewInventoryService crea una nueva instancia del servicio.
func NewInventoryService(
	repository InventoryRepository,
) *InventoryService {

	return &InventoryService{
		repository: repository,
	}
}

// AddInventory registra un nuevo inventario.
func (s *InventoryService) AddInventory(
	inventory InventoryRecord,
) error {

	if inventory.GetStock() < 0 {
		return fmt.Errorf("stock invalido")
	}

	return s.repository.Save(inventory)
}

// MemoryRepository simula una base de datos en memoria.
type MemoryRepository struct {
	items []InventoryRecord
}

// Save almacena un registro de inventario.
func (r *MemoryRepository) Save(
	inventory InventoryRecord,
) error {

	r.items = append(r.items, inventory)

	return nil
}

// GetAll devuelve todos los registros de inventario.
func (r *MemoryRepository) GetAll() (
	[]InventoryRecord,
	error,
) {
	return r.items, nil
}
