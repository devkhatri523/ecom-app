package domain

type Category struct {
	CategoryId          int32  `gorm:"column:id;primary_key;default:nextval('category_id_seq')"`
	CategoryName        string `gorm:"column:name;not null"`
	CategoryDescription string `gorm:"column:description;"`
}

func (c *Category) TableName() string { return "category" }
