package repo

import "payment-service/domain"

type PaymentRepository interface {
	CreatePayment(payment domain.Payment) (int32, error)
}
