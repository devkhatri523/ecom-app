package domain

type OrderLine struct {
	Id        int     `gorm:"column:id;primaryKey"`
	ProductId int     `gorm:"column:product_id"`
	Quantity  float64 `gorm:"column:quantity"`
	OrderId   int32   `gorm:"column:order_id"`
}

func (OrderLine) TableName() string {
	return "order_line"
}
