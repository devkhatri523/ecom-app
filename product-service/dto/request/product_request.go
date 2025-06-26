package request

type ProductRequest struct {
	Name              string  `json:"name" validate:"required"`
	Description       string  `json:"description" validate:"required"`
	AvailableQuantity float64 `json:"availableQuantity" validate:"required,gt=0"`
	Price             float64 `json:"price" validate:"required,gt=0"`
	CategoryId        int32   `json:"categoryId" validate:"required"`
}

type ProductUpdateRequest struct {
	Name              string  `json:"name" validate:"required"`
	Description       string  `json:"description" validate:"required"`
	AvailableQuantity float64 `json:"availableQuantity" validate:"required,gt=0"`
	Price             float64 `json:"price" validate:"required,gt=0"`
}
