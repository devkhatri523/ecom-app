package repo

import "product-service/domain"

type ProductRepository interface {
	CreateProduct(product domain.Product) (domain.Product, error)
	FindAllProducts() ([]domain.Product, error)
	FindProductById(productId int32) (product domain.Product, err error)
	DeleteProduct(productId int32) (int32, error)
	UpdateProduct(productId int32, product domain.Product) (int32, error)
	FindAllProductsByIds(productIds []int32) ([]domain.Product, error)
}
