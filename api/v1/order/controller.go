package loan

import (
	"core-data/api/common"
	"core-data/api/v1/order/request"
	"core-data/api/v1/order/response"
	"core-data/business"
	"core-data/business/order"
	"core-data/util/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	orderService order.Service
}

func NewController(orderService order.Service) *Controller {
	return &Controller{
		orderService,
	}
}

// Insert order
func (controller Controller) InsertOrder(c echo.Context) error {
	Insertorder := new(request.InsertOrderRequest)

	if err := c.Bind(Insertorder); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(Insertorder); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	orderSpec := Insertorder.ToUpsertOrderSpec()
	err := controller.orderService.InsertOrder(orderSpec)
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(business.ErrDuplicate))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

// Update order
func (controller Controller) UpdateOrderStatus(c echo.Context) error {
	updateorder := new(request.UpdateOrderRequest)

	if err := c.Bind(updateorder); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(updateorder); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	orderSpec := updateorder.ToUpdateOrderSpec()
	res, err := controller.orderService.UpdateOrderStatus(*orderSpec)
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponse(response.NewUpdateOrderResponse(*res)))
}
