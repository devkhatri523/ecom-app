package repo

import (
	"order-service/domain"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	Db *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		Db: db,
	}
}

func (o OrderRepositoryImpl) CreateOrder(order domain.Order) (int32, error) {
	result := o.Db.Create(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return order.OrderId, nil
}

func (o OrderRepositoryImpl) FindAllOrders() ([]domain.Order, error) {
	var orders []domain.Order
	result := o.Db.Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
