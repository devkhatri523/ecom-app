package domain

import "payment-service/enum"

type PaymentConfirmation struct {
	OrderReference    string             `json:"orderReference"`
	Amount            float64            `json:"amount"`
	CustomerFirstName string             `json:"customerFirstName"`
	CustomerLastName  string             `json:"customerLastName"`
	PaymentMethod     enum.PaymentMethod `json:"paymentMethod"`
	CustomerEmail     string             `json:"customerEmail"`
}
