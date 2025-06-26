package request

import "payment-service/enum"

type PaymentRequest struct {
	Amount           float64            `json:"amount"`
	PaymentMethod    enum.PaymentMethod `json:"payment_method"`
	OrderId          int32              `json:"orderId"`
	OrderReferenceId string             `json:"orderReferenceId"`
	Customer         `json:"customer"`
}

type Customer struct {
	CustomerId int32  `json:"customerId"`
	Email      string `json:"email"`
}
