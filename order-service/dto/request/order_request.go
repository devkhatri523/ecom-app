package request

import "order-service/enum"

type OrderRequest struct {
	Id               int                    `json:"id"`
	OrderReferenceId string                 `json:"orderReferenceId"`
	Amount           float64                `json:"amount"`
	PaymentMethod    enum.PaymentMethod     `json:"paymentMethod"`
	CustomerId       string                 `json:"customerId"`
	PurchaseProducts []PurchaseOrderRequest `json:"purchaseProducts"`
}

type PurchaseOrderRequest struct {
	ProductId int     `json:"productId"`
	Quantity  float64 `json:"quantity"`
}
