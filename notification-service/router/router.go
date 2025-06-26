package router

import (
	"notification-service/controller"

	"github.com/gin-gonic/gin"
)

func NotificatioRouter(orderNotificationController *controller.OrderNotificationController, paymentNotificationController *controller.PaymentNotificationController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/api/notification")
	router.POST("/orderemail", orderNotificationController.SendOderConfirmationEmail)
	router.POST("/paymentemail", paymentNotificationController.SendPaymentConfirmationEmail)
	return service

}
