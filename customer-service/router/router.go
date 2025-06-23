package router

import (
	"customer-service/controller"

	"github.com/gin-gonic/gin"
)

func CustomerRouter(customerController *controller.CustomerController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/api/customer")
	router.GET("", customerController.GetAllCustomers)
	router.POST("", customerController.CreateCustomer)
	router.GET("/:customerId", customerController.GetCustomerById)
	router.PATCH("/:customerId", customerController.UpdateCustomer)
	router.DELETE("/:customerId", customerController.DeleteCustomer)
	router.GET("/cutomerexists/:customerId", customerController.IsCustomerExists)
	return service
}
