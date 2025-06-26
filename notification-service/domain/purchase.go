package domain

type Purchase struct {
	ProductId   int     `json:"productId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
}
