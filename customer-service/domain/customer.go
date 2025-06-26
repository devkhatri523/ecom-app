package domain

type Customer struct {
	Id        int64   `gorm:"column:id;primary_key;default:nextval('customer_id_seq')"`
	FirstName string  `gorm:"column:first_name;not null"`
	LastName  string  `gorm:"column:last_name;not null"`
	Email     string  `gorm:"column:email;not null"`
	Address   Address `gorm:"embedded"`
}

func (c Customer) TableName() string { return "customer" }
