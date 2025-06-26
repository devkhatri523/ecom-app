package service

import (
	"product-service/domain"
	"product-service/dto/request"
	"product-service/dto/response"
)

type ProductService interface {
	CreateProduct(request request.ProductRequest) (domain.Product, error)
	FindAll() ([]response.ProductResponse, error)
	FindById(productId int32) (response.ProductResponse, error)
	Delete(productId int32) (int32, error)
	Update(productId int32, request request.ProductUpdateRequest) (int32, error)
	PurchaseProducts(purchaseProducts []request.ProductPurchaseRequest) ([]response.ProductPurchaseResponse, error)
}
