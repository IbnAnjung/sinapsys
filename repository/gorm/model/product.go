package model

type Product struct {
	ID          uint64  `gorm:"column:id"`
	Name        string  `gorm:"column:name"`
	CategoryID  uint64  `gorm:"column:category_id"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
}

func (m *Product) TableName() string {
	return "products"
}

// GetProductLists
type GetProductList struct {
	ID                uint64  `gorm:"column:id"`
	ProductID         uint64  `gorm:"column:product_id"`
	Name              string  `gorm:"column:name"`
	ProductCategoryID uint64  `gorm:"column:category_id"`
	Description       string  `gorm:"column:description"`
	Price             float64 `gorm:"column:price"`
	CategoryName      string  `gorm:"column:category_name"`
}
