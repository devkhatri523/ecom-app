package service

import "payment-service/dto/request"

type CustomerService interface {
	CreatePayment(request request.PaymentRequest) (int32, error)
}
