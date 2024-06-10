package entity

import (
	"context"
	"time"

	paymentmanualtransfer "github.com/IbnAnjung/synapsis/entity/dto/payment_manual_transfer"
)

type ManualTransferPaymentStatus uint8

const (
	ManualTransferPaymentStatusCreated ManualTransferPaymentStatus = 1
	ManualTransferPaymentStatusInvalid ManualTransferPaymentStatus = 9
	ManualTransferPaymentStatusValid   ManualTransferPaymentStatus = 10
)

type PaymentManualTransfer struct {
	ID                uint64
	OrderPaymentID    uint64
	BankAccountNumber string
	BankAccountName   string
	Date              time.Time
	Value             float64
	Status            ManualTransferPaymentStatus
}

type PaymentUseCase interface {
	ManualTransferConfirmation(ctx context.Context, input paymentmanualtransfer.PaymentManualTransferConfirmation) (err error)
}

type PaymentManualTransferRepository interface {
	Create(ctx context.Context, payment *PaymentManualTransfer) (err error)
}
