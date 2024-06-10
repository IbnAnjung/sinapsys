package model

import (
	"time"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/jinzhu/copier"
)

type PaymentManualTransfer struct {
	ID                uint64    `gorm:"column:id;primaryKey;auto_increment"`
	OrderPaymentID    uint64    `gorm:"column:order_payment_id"`
	BankAccountNumber string    `gorm:"column:bank_account_number"`
	BankAccountName   string    `gorm:"column:bank_account_name"`
	Date              time.Time `gorm:"column:date"`
	Value             float64   `gorm:"column:value"`
	Status            uint8     `gorm:"column:status"`
}

func (m *PaymentManualTransfer) TableName() string {
	return "payment_manual_transfers"
}

func (m *PaymentManualTransfer) ToEntity() (en entity.PaymentManualTransfer) {
	copier.Copy(&en, &m)
	en.Status = entity.ManualTransferPaymentStatus(m.Status)
	return
}

func (m *PaymentManualTransfer) FillFromEntity(en entity.PaymentManualTransfer) {
	copier.Copy(m, &en)
	m.Status = uint8(en.Status)
}
