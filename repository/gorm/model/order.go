package model

import (
	"time"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/jinzhu/copier"
)

type Order struct {
	ID          uint64    `gorm:"column:id;primaryKey;auto_increment"`
	UserID      uint64    `gorm:"column:user_id"`
	CreatedDate time.Time `gorm:"column:created_date"`
	ExpiredDate time.Time `gorm:"column:expired_date"`
	TotalPrice  float64   `gorm:"column:total_price"`
	Status      uint8     `gorm:"column:status"`
}

func (m *Order) TableName() string {
	return "orders"
}

func (m *Order) ToEntity() (en entity.Order) {
	copier.Copy(&en, &m)
	en.Status = entity.OrderStatus(m.Status)
	return
}

func (m *Order) FillFromEntity(en entity.Order) {
	copier.Copy(m, &en)
	m.Status = uint8(en.Status)
}
