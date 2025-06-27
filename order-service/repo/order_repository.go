package repo

import "order-service/domain"

type OrderRepository interface {
	CreateOrder(order domain.Order) (int32, error)
	FindAllOrders() ([]domain.Order, error)
}
