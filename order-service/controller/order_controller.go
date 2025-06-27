package controller

import (
	"net/http"
	"order-service/dto/request"
	"order-service/dto/response"
	"order-service/service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (o *OrderController) FindAllOrder(ctx *gin.Context) {
	data, err := o.orderService.FindAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	req := request.OrderRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	orderId, err := o.orderService.CreateOrder(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   orderId,
	}
	ctx.JSON(http.StatusOK, res)
}
