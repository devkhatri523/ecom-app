package controller

import (
	"net/http"
	"notification-service/domain"
	service "notification-service/emailservice"

	"github.com/gin-gonic/gin"
)

type OrderNotificationController struct {
	service service.EmailService
}

func NewOrderNotificationController(service service.EmailService) *NotificationController {
	return &NotificationController{service: service}
}

func (controller *OrderNotificationController) SendOderConfirmationEmail(ctx *gin.Context) {
	req := domain.OrderConfirmation{}
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
