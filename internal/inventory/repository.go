package inventory

type InventoryRepository interface {
	Save(inventory InventoryRecord) error
	GetAll() ([]InventoryRecord, error)
}
