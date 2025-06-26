package router

import (
	"payment-service/controller"

	"github.com/gin-gonic/gin"
)

func PaymentRouter(paymentController *controller.PaymentController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/api/payment")
	router.POST("", paymentController.CreatePayment)
	return service

}
