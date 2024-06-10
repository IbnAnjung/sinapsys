package model

import (
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/jinzhu/copier"
)

type OrderProduct struct {
	ID          uint64  `gorm:"column:id;primaryKey;auto_increment"`
	OrderID     uint64  `gorm:"column:order_id"`
	ProductID   uint64  `gorm:"column:product_id"`
	ProductName string  `gorm:"column:product_name"`
	Quantity    float64 `gorm:"column:quantity"`
	Price       float64 `gorm:"column:price"`
}

func (m *OrderProduct) TableName() string {
	return "order_products"
}

func (m *OrderProduct) ToEntity() (en entity.OrderProduct) {
	copier.Copy(&en, &m)
	return
}

func (m *OrderProduct) FillFromEntity(en entity.OrderProduct) {
	copier.Copy(&m, &en)
}
