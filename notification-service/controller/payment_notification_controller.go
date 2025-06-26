package controller

import (
	"net/http"
	"notification-service/domain"
	service "notification-service/emailservice"

	"github.com/gin-gonic/gin"
)

type PaymentNotificationController struct {
	service service.EmailService
}

func NewPaymentNotificationController(service service.EmailService) *PaymentNotificationController {
	return &PaymentNotificationController{service: service}
}

func (controller *PaymentNotificationController) SendPaymentConfirmationEmail(ctx *gin.Context) {
	req := domain.PaymentConfirmation{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

	}
	err = controller.service.SendEmail(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

}
