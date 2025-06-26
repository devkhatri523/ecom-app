package router

import (
	"product-service/controller"

	"github.com/gin-gonic/gin"
)

func ProductRouter(productController *controller.ProductController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/api/products")
	router.GET("", productController.GetAllProducts)
	router.POST("", productController.CreateProduct)
	router.DELETE("/:productId", productController.DeleteProduct)
	router.PATCH("/:productId", productController.UpdateProduct)
	router.GET("/:productId", productController.GetProduct)
	router.POST("/purchase", productController.PuchaseProduct)
	return service

}
