package service

import (
	"log"
	"order-service/client"
	"order-service/domain"
	"order-service/dto/request"
	"order-service/dto/response"
	"order-service/repo"
	"strconv"
	"time"
)

type OrderServiceImpl struct {
	orderRepository    repo.OrderRepository
	orderLinRepository repo.OrderLineRepository
}

func NewOrderServiceImpl(orderRepository repo.OrderRepository, orderLinRepository repo.OrderLineRepository) OrderService {
	return &OrderServiceImpl{
		orderRepository:    orderRepository,
		orderLinRepository: orderLinRepository,
	}
}

func (o OrderServiceImpl) CreateOrder(req request.OrderRequest) (orderId int32, err error) {
	customer, err := client.GetCustomerDetails(req.CustomerId)
	if err != nil {
		return 0, err
	}
	custId, _ := strconv.Atoi(customer.CustomerID)
	purchaseProducts, err := client.PurchaseProducts(req)
	if err != nil {
		return 0, err
	}
	order := domain.Order{
		CustomerId:       req.CustomerId,
		OrderReference:   req.OrderReferenceId,
		TotalAmount:      req.Amount,
		PaymentMethod:    req.PaymentMethod.String(),
		CreatedAt:        time.Now(),
		LastModifiedDate: time.Now(),
	}

	orderId, err = o.orderRepository.CreateOrder(order)
	if err != nil {
		return 0, err
	}
	for _, value := range purchaseProducts {
		purchaseProduct := domain.OrderLine{
			ProductId: int(value.ProductId),
			Quantity:  value.Quantity,
			OrderId:   orderId,
		}
		err = o.orderLinRepository.SaveOrderLine(purchaseProduct)
		if err != nil {
			return 0, err
		}
	}

	paymentRequest := request.PaymentRequest{
		Amount:           req.Amount,
		PaymentMethod:    req.PaymentMethod,
		OrderReferenceId: req.OrderReferenceId,
		OrderId:          order.OrderId,
		Customer: request.Customer{
			CustomerId: int32(custId),
			Email:      customer.Email,
		},
	}
	err = client.RequestPayment(paymentRequest)
	if err != nil {
		log.Fatal(err)

	}
	var purchaseOrderList []request.OrderConfirmationPurchaseRequest

	for _, value := range purchaseProducts {
		orderConfirmationPurchaseReq := request.OrderConfirmationPurchaseRequest{
			ProductId:   int(value.ProductId),
			Name:        value.ProductName,
			Description: value.Description,
			Price:       value.Price,
			Quantity:    value.Quantity,
		}
		purchaseOrderList = append(purchaseOrderList, orderConfirmationPurchaseReq)
	}
	orderConfirmationReuest := request.OrderConfirmation{
		OrderReferenceNumber: req.OrderReferenceId,
		TotalAmount:          req.Amount,
		PaymentMethod:        req.PaymentMethod,
		Customer: request.OrderConfirmationCustomerRequest{
			Id:        custId,
			Email:     customer.Email,
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
		},
		Products: purchaseOrderList,
	}
	err = client.SendOrderConfirmation(orderConfirmationReuest)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return order.OrderId, nil
}

func (o OrderServiceImpl) FindAllOrders() (orders []response.OrderResponse, err error) {
	var orderReponse []response.OrderResponse
	result, err := o.orderRepository.FindAllOrders()
	if err != nil {
		return nil, err
	}
	for _, value := range result {
		order := response.OrderResponse{
			Id:              value.OrderId,
			CustomerId:      value.CustomerId,
			Amount:          value.TotalAmount,
			PaymentMethod:   value.PaymentMethod,
			ReferenceNumber: value.OrderReference,
		}
		orderReponse = append(orderReponse, order)
	}
	return orderReponse, nil

}
