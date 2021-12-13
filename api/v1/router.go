package router

import (
	orderController "core-data/api/v1/order"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	orderController orderController.Controller,
) {

	// order
	orderV1 := e.Group("v1/order")
	orderV1.POST("/add", orderController.InsertOrder)
	orderV1.PUT("/update-status", orderController.UpdateOrderStatus)
	// orderV1.GET("", orderController.GetAllorder)

}
