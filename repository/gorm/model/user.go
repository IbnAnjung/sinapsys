package model

import (
	"github.com/IbnAnjung/synapsis/entity"
)

type MUser struct {
	ID          int64  `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name"`
	PhoneNumber string `gorm:"column:phone_number"`
	Password    string `gorm:"column:password"`
}

func (m *MUser) TableName() string {
	return "users"
}

func (m *MUser) ToEntity() (en entity.User) {
	en.ID = m.ID
	en.Name = m.Name
	en.PhoneNumber = m.PhoneNumber
	en.Password = m.Password
	return
}

func (m *MUser) FillFromEntity(en entity.User) {
	m.ID = en.ID
	m.Name = en.Name
	m.PhoneNumber = en.PhoneNumber
	m.Password = en.Password
}
