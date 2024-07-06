package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-order-stock-manager/model"
	"simple-order-stock-manager/utils"
)

func (c *Config) placeOrder(context *gin.Context) {

	var body model.PlaceOrderHTTPRequest
	bindingError := context.ShouldBindJSON(&body)
	if bindingError != nil {
		c.context.Logger().Error("Error while binding request, error: ", bindingError.Error())
		context.JSON(http.StatusBadRequest, model.PlaceOrderResponse{
			Message: model.InvalidRequest,
		})
		return
	}

	request, validationError := c.validate(body)
	if validationError != nil {
		context.JSON(http.StatusBadRequest, model.PlaceOrderResponse{
			Message: model.InvalidRequest,
		})
		return
	}

	orderId, placingOrderError := c.services.PlaceOrderProcessing(*request)
	if placingOrderError != nil {
		c.context.Logger().Error("PlaceOrderProcessing error, error: ", placingOrderError.Error())
		if utils.InArrayError(placingOrderError, model.PlacingOrderErrors) {
			context.JSON(http.StatusBadRequest, model.PlaceOrderResponse{
				Message: placingOrderError.Error(),
			})
			return
		}
		context.JSON(http.StatusInternalServerError, model.PlaceOrderResponse{
			Message: "Internal server error",
		})
		return
	}

	context.JSON(http.StatusOK, model.PlaceOrderResponse{
		Message: model.OrderPlacedSuccessfully,
		Id:      orderId,
	})
	return
}
