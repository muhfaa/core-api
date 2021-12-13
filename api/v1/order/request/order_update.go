package request

import (
	"core-data/business/order"
)

type UpdateOrderRequest struct {
	ID      int `json:"id" validate:"required"`
	Version int `json:"version" validate:"required"`
}

func (req *UpdateOrderRequest) ToUpdateOrderSpec() *order.UpdateOrderStatusSpec {
	var updateSpec order.UpdateOrderStatusSpec
	updateSpec.ID = req.ID
	updateSpec.Version = req.Version

	return &updateSpec
}
