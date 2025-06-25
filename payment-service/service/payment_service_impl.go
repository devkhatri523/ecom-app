package service

import (
	"payment-service/domain"
	"payment-service/dto/request"
	"payment-service/repo"
	"time"
)

type PaymentServiceImpl struct {
	paymentRepositoty repo.PaymentRepository
}

func NewCustomerServiceImpl(paymentRepository repo.PaymentRepository) PaymentService {
	return &PaymentServiceImpl{
		paymentRepositoty: paymentRepository,
	}
}

func (p PaymentServiceImpl) CreatePayment(request request.PaymentRequest) (int32, error) {
	payment := domain.Payment{
		Amount:        request.Amount,
		CustomerId:    request.Customer.CustomerId,
		OrderId:       request.OrderId,
		PaymentMethod: request.PaymentMethod.String(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	paymentId, err := p.paymentRepositoty.CreatePayment(payment)
	if err != nil {
		return 0, err
	}
	return paymentId, nil
}
