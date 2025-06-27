package service

import (
	"order-service/dto/request"
	"order-service/dto/response"
)

type OrderService interface {
	CreateOrder(req request.OrderRequest) (orderId int32, err error)
	FindAllOrders() (orders []response.OrderResponse, err error)
}
