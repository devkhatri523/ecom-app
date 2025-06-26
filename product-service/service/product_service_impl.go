package service

import (
	"errors"
	"fmt"
	"product-service/domain"
	"product-service/dto/request"
	"product-service/dto/response"
	"product-service/mapper"
	"product-service/repo"
	"sort"
)

type ProductServiceImpl struct {
	productRepository repo.ProductRepository
}

func NewProductServiceImpl(productRepository repo.ProductRepository) ProductService {
	return &ProductServiceImpl{
		productRepository: productRepository,
	}
}

func (p ProductServiceImpl) CreateProduct(request request.ProductRequest) (domain.Product, error) {
	product := domain.Product{
		ProductName:        request.Name,
		ProductDescription: request.Description,
		AvailableQuantity:  request.AvailableQuantity,
		Price:              request.Price,
		CategoryId:         request.CategoryId,
		Category: domain.Category{
			CategoryId: request.CategoryId,
		},
	}
	result, err := p.productRepository.CreateProduct(product)
	if err != nil {
		return domain.Product{}, err
	}
	return result, nil
}

func (p ProductServiceImpl) FindAll() ([]response.ProductResponse, error) {
	var productResponse []response.ProductResponse
	result, err := p.productRepository.FindAllProducts()
	if err != nil {
		return nil, err
	}
	for _, value := range result {
		productRes := response.ProductResponse{
			Id:                 value.Id,
			ProductName:        value.ProductName,
			ProductDescription: value.ProductDescription,
			AvailableQuantity:  value.AvailableQuantity,
			Price:              value.Price,
			Category: response.ProductCategory{
				CategoryId:          value.CategoryId,
				CategoryName:        value.Category.CategoryName,
				CategoryDescription: value.Category.CategoryDescription,
			},
		}
		productResponse = append(productResponse, productRes)
	}
	return productResponse, nil
}

func (p ProductServiceImpl) FindById(productId int32) (response.ProductResponse, error) {
	data, err := p.productRepository.FindProductById(productId)
	if err != nil {
		return response.ProductResponse{}, err
	}
	res := response.ProductResponse{
		Id:                 data.Id,
		ProductName:        data.ProductName,
		ProductDescription: data.ProductDescription,
		AvailableQuantity:  data.AvailableQuantity,
		Price:              data.Price,
		Category: response.ProductCategory{
			CategoryId:          data.CategoryId,
			CategoryName:        data.Category.CategoryName,
			CategoryDescription: data.Category.CategoryDescription,
		},
	}
	return res, nil
}

func (p ProductServiceImpl) Delete(productId int32) (int32, error) {
	productId, err := p.productRepository.DeleteProduct(productId)
	if err != nil {
		return 0, err
	}
	return productId, nil
}

func (p ProductServiceImpl) Update(productId int32, request request.ProductUpdateRequest) (int32, error) {
	product := domain.Product{
		ProductName:        request.Name,
		ProductDescription: request.Description,
		AvailableQuantity:  request.AvailableQuantity,
		Price:              request.Price,
	}
	productId, err := p.productRepository.UpdateProduct(productId, product)
	if err != nil {
		return 0, err
	}
	return productId, nil
}

func (p ProductServiceImpl) PurchaseProducts(purchaseProductRequest []request.ProductPurchaseRequest) ([]response.ProductPurchaseResponse, error) {
	var productIds []int32
	var productPurchaseResponses []response.ProductPurchaseResponse
	for _, value := range purchaseProductRequest {
		productIds = append(productIds, value.Id)
	}
	storedProducts, err := p.productRepository.FindAllProductsByIds(productIds)
	if err != nil {
		return nil, err
	}
	if len(storedProducts) != len(productIds) {
		return nil, errors.New("One or more product does not exits")
	}
	sort.Slice(purchaseProductRequest, func(i, j int) bool {
		return purchaseProductRequest[i].Id > purchaseProductRequest[j].Id

	})
	for i := 0; i < len(storedProducts); i++ {
		product := storedProducts[i]
		productRequest := purchaseProductRequest[i]
		if product.AvailableQuantity < productRequest.Quantity {
			return nil, errors.New(fmt.Sprintf("Insufficient stock for product with id :%d", product.Id))
		}

		newAvailableQuantity := product.AvailableQuantity - productRequest.Quantity
		product.AvailableQuantity = newAvailableQuantity
		_, er := p.productRepository.UpdateProduct(product.Id, product)
		if er != nil {
			return nil, errors.New("Exception occured while updating product")
		}
		productPurchaseResponses = append(productPurchaseResponses, mapper.MapToPurchaseResponse(product, product.AvailableQuantity))

	}
	return productPurchaseResponses, nil

}
