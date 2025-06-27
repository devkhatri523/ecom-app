package domain

import "time"

type Order struct {
	OrderId          int32     `gorm:"column:id;primaryKey"`
	OrderReference   string    `gorm:"column:reference"`
	TotalAmount      float64   `gorm:"column:total_amount"`
	PaymentMethod    string    `gorm:"column:payment_method"`
	CustomerId       string    `gorm:"column:customer_id"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	LastModifiedDate time.Time `gorm:"column:last_modified_date"`
}

func (Order) TableName() string {
	return "customer_order"
}
