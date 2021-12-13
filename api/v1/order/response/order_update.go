package response

import (
	"core-data/business/order"
)

type UpdateOrderResponse struct {
	ID            int    `json:"id"`
	StatusService string `json:"status_service"`
}

func NewUpdateOrderResponse(res order.UpdateOrderStatus) UpdateOrderResponse {
	var updateSpec UpdateOrderResponse
	updateSpec.ID = res.ID
	updateSpec.StatusService = string(res.StatusService)

	return updateSpec
}
