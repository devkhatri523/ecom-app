package service

import "payment-service/dto/request"

type PaymentService interface {
	CreatePayment(request request.PaymentRequest) (int32, error)
}
