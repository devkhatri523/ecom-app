package request

type ProductPurchaseRequest struct {
	Id       int32   `json:"productId" validate:"required"`
	Quantity float64 `json:"quantity" validate:"required,gt=0"`
}
