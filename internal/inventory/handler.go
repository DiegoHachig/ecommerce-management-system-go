package inventory

import "fmt"

// Inventory represents an item in inventory with methods used by handlers.
type Inventory interface {
	GetProductID() string
	GetStock() int
}

func ShowInventory(
	inventories []InventoryRecord,
) {

	for _, item := range inventories {

		fmt.Println(
			item.GetProductID(),
			item.GetStock(),
		)
	}
}
