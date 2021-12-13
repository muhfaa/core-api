package order

type Repository interface {
	InsertOrder(insertSpec OrderSpec) error
	UpdateOrderStatus(updateSpec UpdateOrderStatus) (bool, error)
	FindOrderByID(id int) (*Order, error)
}
