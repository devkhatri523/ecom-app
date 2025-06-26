package controller

import (
	"net/http"
	"payment-service/dto/request"
	"payment-service/dto/response"
	"payment-service/service"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}

}

func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	req := request.PaymentRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return

	}
	paymentId, err := c.paymentService.CreatePayment(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   paymentId,
	}
	ctx.JSON(http.StatusOK, res)
}
