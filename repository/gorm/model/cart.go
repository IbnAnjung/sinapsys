package model

type Cart struct {
	ID        uint64  `gorm:"column:id;primaryKey;auto_increment"`
	UserID    uint64  `gorm:"column:user_id"`
	ProductID uint64  `gorm:"column:product_id"`
	Quantity  float64 `gorm:"column:quantity"`
}

func (m *Cart) TableName() string {
	return "carts"
}

// GetCart
type GetCart struct {
	ID         uint64  `gorm:"column:id"`
	ProductID  uint64  `gorm:"column:product_id"`
	Name       string  `gorm:"column:name"`
	Quantity   string  `gorm:"column:quantity"`
	Price      float64 `gorm:"column:price"`
	TotalPrice float64 `gorm:"column:total_price"`
}
