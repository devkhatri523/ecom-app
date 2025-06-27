package response

type OrderResponse struct {
	Id              int32   `json:"id"`
	ReferenceNumber string  `json:"referenceNumber"`
	Amount          float64 `json:"amount"`
	PaymentMethod   string  `json:"paymentMethod"`
	CustomerId      string  `json:"customerId"`
}
