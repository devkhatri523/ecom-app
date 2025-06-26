package domain

import "time"

type Payment struct {
	Id            int32     `gorm:"column:id;primaryKey"`
	Amount        float64   `gorm:"column:amount"`
	PaymentMethod string    `gorm:"column:payment_method"`
	OrderId       int32     `gorm:"column:order_id"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	CustomerId    int32     `gorm:"column:customer_id"`
}

func (Payment) TableName() string {
	return "payment"
}
