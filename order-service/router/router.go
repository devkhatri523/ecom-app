package router

import (
	"order-service/controller"

	"github.com/gin-gonic/gin"
)

func OrderRouter(controller *controller.OrderController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/api/order")
	router.GET("", controller.FindAllOrder)
	router.POST("", controller.CreateOrder)
	return service

}
