package orders

import "errors"

// Order representa un pedido.
type Order struct {
	id     int
	userID int
	total  float64
	status string
}

// Constructor
func NewOrder(
	id int,
	userID int,
	total float64,
	status string,
) (*Order, error) {

	if total <= 0 {
		return nil, errors.New("el total debe ser mayor a cero")
	}

	return &Order{
		id:     id,
		userID: userID,
		total:  total,
		status: status,
	}, nil
}

func (o Order) GetID() int {
	return o.id
}

func (o Order) GetUserID() int {
	return o.userID
}

func (o Order) GetTotal() float64 {
	return o.total
}

func (o Order) GetStatus() string {
	return o.status
}
