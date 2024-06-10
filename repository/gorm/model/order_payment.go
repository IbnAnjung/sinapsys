package model

import (
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/jinzhu/copier"
)

type OrderPayment struct {
	ID      uint64  `gorm:"column:id;primaryKey;auto_increment"`
	OrderID uint64  `gorm:"column:order_id"`
	Value   float64 `gorm:"column:value"`
	Type    uint8   `gorm:"column:type"`
	Status  uint8   `gorm:"column:status"`
}

func (m *OrderPayment) TableName() string {
	return "order_payments"
}

func (m *OrderPayment) ToEntity() (en entity.OrderPayment) {
	copier.Copy(&en, &m)
	en.Status = entity.OrderPaymentStatus(m.Status)
	return
}

func (m *OrderPayment) FillFromEntity(en entity.OrderPayment) {
	copier.Copy(m, &en)
	m.Status = uint8(en.Status)
}
