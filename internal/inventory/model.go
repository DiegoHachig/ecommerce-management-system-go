package inventory

import "errors"

// InventoryRecord representa el inventario de un producto.
type InventoryRecord struct {
	productID int
	stock     int
}

// NewInventory crea un nuevo registro de inventario.
func NewInventory(
	productID int,
	stock int,
) (*InventoryRecord, error) {

	if stock < 0 {
		return nil, errors.New("stock invalido")
	}

	return &InventoryRecord{
		productID: productID,
		stock:     stock,
	}, nil
}

// GetProductID devuelve el identificador del producto.
func (i InventoryRecord) GetProductID() int {
	return i.productID
}

// GetStock devuelve el stock disponible.
func (i InventoryRecord) GetStock() int {
	return i.stock
}
