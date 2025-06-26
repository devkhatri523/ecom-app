package domain

import "notification-service/enum"

type OrderConfirmation struct {
	OrderReferenceNumber string             `json:"orderReferenceNumber"`
	TotalAmount          float64            `json:"totalAmount"`
	PaymentMethod        enum.PaymentMethod `json:"paymentMethod"`
	Customer             Customer           `json:"customer"`
	Products             []Purchase         `json:"products"`
}
