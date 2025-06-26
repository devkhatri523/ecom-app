package response

type ProductResponse struct {
	Id                 int32           `json:"productId"`
	ProductName        string          `json:"productName" `
	ProductDescription string          `json:"productDescription" `
	AvailableQuantity  float64         `json:"availableQuantity"`
	Price              float64         `json:"price"`
	Category           ProductCategory `json:"category"`
}

type ProductCategory struct {
	CategoryId          int32  `json:"categoryId"`
	CategoryName        string `json:"categoryName"`
	CategoryDescription string `json:"categoryDescription"`
}
