package response

type ProductPurchaseResponse struct {
	ProductId   int32   `json:"productId"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
}
