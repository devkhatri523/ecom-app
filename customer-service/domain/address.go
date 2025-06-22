package domain

type Address struct {
	Street      string `gorm:"column:street;not null"`
	HouseNumber string `gorm:"column:house_number;not null"`
	ZipCode     string `gorm:"column:zip_code;not null"`
}
