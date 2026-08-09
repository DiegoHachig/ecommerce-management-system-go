package orders

import "fmt"

// ShowOrders muestra los pedidos.
func ShowOrders(
	orders []Order,
) {

	for _, order := range orders {

		fmt.Println(
			order.GetID(),
			order.GetUserID(),
			order.GetTotal(),
			order.GetStatus(),
		)
	}
}
