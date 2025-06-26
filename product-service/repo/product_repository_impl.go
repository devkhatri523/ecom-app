package repo

import (
	"errors"
	"fmt"
	"product-service/domain"
	"product-service/helper"
	"product-service/query"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProdcuRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		Db: Db,
	}

}

func (p ProductRepositoryImpl) CreateProduct(product domain.Product) (domain.Product, error) {
	result := p.Db.Create(&product)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	return product, nil
}

func (p ProductRepositoryImpl) FindAllProducts() ([]domain.Product, error) {
	var productCategory []domain.ProductCategory
	var products []domain.Product
	stmt := p.Db.Raw(query.PRODUCT_QUERY)
	results := stmt.Find(&productCategory)
	if results.Error != nil {
		return nil, results.Error
	}
	for _, value := range productCategory {
		product := domain.Product{
			Id:                 value.Id,
			ProductName:        value.ProductName,
			ProductDescription: value.ProductDescription,
			AvailableQuantity:  value.AvailableQuantity,
			Price:              value.Price,
			Category: domain.Category{
				CategoryId:          value.CategoryId,
				CategoryName:        value.CategoryName,
				CategoryDescription: value.CategoryDescription,
			},
		}
		products = append(products, product)
	}
	return products, nil

}

func (p ProductRepositoryImpl) FindProductById(productId int32) (product domain.Product, err error) {
	var indvProductCategoty domain.ProductCategory
	whereQuery := fmt.Sprintf("%s where p.id=%d", query.PRODUCT_QUERY, productId)
	stmt := p.Db.Raw(whereQuery)
	result := stmt.Find(&indvProductCategoty)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}
	if result.RowsAffected == 0 {
		return domain.Product{}, errors.New("product not found")
	}
	indvProduct := domain.Product{
		Id:                 indvProductCategoty.Id,
		ProductName:        indvProductCategoty.ProductName,
		ProductDescription: indvProductCategoty.ProductDescription,
		AvailableQuantity:  indvProductCategoty.AvailableQuantity,
		Price:              indvProductCategoty.Price,
		Category: domain.Category{
			CategoryId:          indvProductCategoty.CategoryId,
			CategoryName:        indvProductCategoty.CategoryName,
			CategoryDescription: indvProductCategoty.CategoryDescription,
		},
	}
	return indvProduct, nil
}

func (p ProductRepositoryImpl) DeleteProduct(productId int32) (int32, error) {
	var product domain.Product
	result := p.Db.Where("id=?", productId).Delete(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return productId, nil
}

func (p ProductRepositoryImpl) UpdateProduct(productId int32, product domain.Product) (int32, error) {
	result := p.Db.Where(&domain.Product{Id: productId}).Updates(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return productId, nil
}

func (p ProductRepositoryImpl) FindAllProductsByIds(productIds []int32) ([]domain.Product, error) {
	var products []domain.Product
	whereQuery := fmt.Sprintf("%s where p.id in (%s) ", query.PRODUCT_QUERY, helper.BuildInCluaseQuery(productIds))
	smt := p.Db.Raw(whereQuery)
	results := smt.Find(&products)
	if results.Error != nil {
		return nil, results.Error
	}
	if results.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}
	return products, nil

}
