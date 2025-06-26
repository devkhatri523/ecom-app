package mapper

import (
	"product-service/domain"
	"product-service/dto/response"
)

func MapToPurchaseResponse(product domain.Product, quantity float64) response.ProductPurchaseResponse {
	productPurchaseReponse := response.ProductPurchaseResponse{
		ProductId:   product.Id,
		ProductName: product.ProductName,
		Description: product.ProductDescription,
		Price:       product.Price,
		Quantity:    quantity,
	}
	return productPurchaseReponse

}
