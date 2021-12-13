package request

import "core-data/business/order"

type InsertOrderRequest struct {
	KerusakanID   int    `json:"kerusakan_id" validate:"required"`
	JenisHP       string `json:"jenis_hp" validate:"required"`
	JenisPlatform string `json:"jenis_platform" validate:"required"`
}

func (req *InsertOrderRequest) ToUpsertOrderSpec() order.OrderSpecRequest {
	var insertSpec order.OrderSpecRequest
	insertSpec.KerusakanID = req.KerusakanID
	insertSpec.JenisHP = req.JenisHP
	insertSpec.JenisPlatform = req.JenisPlatform

	return insertSpec
}
