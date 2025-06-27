package request

import "order-service/enum"

type OrderConfirmation struct {
	OrderReferenceNumber string                             `json:"orderReferenceNumber"`
	TotalAmount          float64                            `json:"totalAmount"`
	PaymentMethod        enum.PaymentMethod                 `json:"paymentMethod"`
	Customer             OrderConfirmationCustomerRequest   `json:"customer"`
	Products             []OrderConfirmationPurchaseRequest `json:"products"`
}

type OrderConfirmationPurchaseRequest struct {
	ProductId   int     `json:"productId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
}

type OrderConfirmationCustomerRequest struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
