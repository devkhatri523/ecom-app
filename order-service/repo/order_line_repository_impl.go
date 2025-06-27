package repo

import (
	"order-service/domain"

	"gorm.io/gorm"
)

type OrderLineRepositoryImpl struct {
	Db *gorm.DB
}

func NewOrderLineRepositoryImpl(db *gorm.DB) OrderLineRepository {
	return &OrderLineRepositoryImpl{
		Db: db,
	}
}

func (o OrderLineRepositoryImpl) SaveOrderLine(orderLine domain.OrderLine) error {
	result := o.Db.Create(&orderLine)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
