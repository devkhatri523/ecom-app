package service

import (
	"payment-service/client"
	"payment-service/domain"
	"payment-service/dto/request"
	"payment-service/repo"
	"strconv"
	"time"
)

type PaymentServiceImpl struct {
	paymentRepositoty repo.PaymentRepository
}

func NewPaymentServiceImpl(paymentRepository repo.PaymentRepository) PaymentService {
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

	customer, _ := client.GetCustomerDetails(strconv.Itoa(int(request.Customer.CustomerId)))
	req := domain.PaymentConfirmation{
		CustomerEmail:     customer.Email,
		CustomerFirstName: customer.FirstName,
		CustomerLastName:  customer.LastName,
		PaymentMethod:     request.PaymentMethod,
		OrderReference:    request.OrderReferenceId,
		Amount:            request.Amount,
	}
	err = client.SendPaymentConfirmation(req)
	if err != nil {
		return paymentId, err
	}

	if err != nil {
		return 0, err
	}
	return paymentId, nil
}
