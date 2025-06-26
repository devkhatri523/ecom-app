package controller

import (
	"fmt"
	"net/http"
	"product-service/dto/request"
	"product-service/dto/response"
	"product-service/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) GetAllProducts(ctx *gin.Context) {
	data, err := productController.productService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   http.StatusOK,
		Data:   data,
		Status: "OK",
	}
	ctx.JSON(http.StatusOK, res)
}

func (productController *ProductController) CreateProduct(ctx *gin.Context) {
	req := request.ProductRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	product, err := productController.productService.CreateProduct(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}
	ctx.JSON(http.StatusOK, res)

}
func (productController *ProductController) GetProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return

	}
	product, err := productController.productService.FindById(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   product,
	}
	ctx.JSON(http.StatusOK, res)
}

func (productController *ProductController) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return

	}
	productRes, err := productController.productService.FindById(int32(id))
	if productRes.Id <= 0 {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Product not found",
		})
		return
	}
	prodId, err := productController.productService.Delete(int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   fmt.Sprintf("Product with id :%d deleted sucessfully", prodId),
	}
	ctx.JSON(http.StatusOK, res)
}
func (productController *ProductController) UpdateProduct(ctx *gin.Context) {
	productUpdateRequest := request.ProductUpdateRequest{}
	err := ctx.ShouldBindJSON(&productUpdateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	updateProductId := ctx.Param("productId")
	id, err := strconv.Atoi(updateProductId)
	productRes, err := productController.productService.FindById(int32(id))
	if productRes.Id <= 0 {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Product not found",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	prodId, err := productController.productService.Update(int32(id), productUpdateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   fmt.Sprintf("Product with id :%d updated sucessfully", prodId),
	}
	ctx.JSON(http.StatusOK, res)
}

func (productController *ProductController) PuchaseProduct(ctx *gin.Context) {
	var purchaseProductRequest []request.ProductPurchaseRequest
	err := ctx.ShouldBindJSON(&purchaseProductRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	purchaseProducts, err := productController.productService.PurchaseProducts(purchaseProductRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   purchaseProducts,
	}
	ctx.JSON(http.StatusOK, res)

}
