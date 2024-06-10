package entity

import (
	"context"
)

type PaymentType uint8
type OrderPaymentStatus uint8

const (
	PaymentTypeManualTransfer PaymentType = 1

	OrderPaymentStatusCreated   OrderPaymentStatus = 1
	OrderPaymentStatusCancelled OrderPaymentStatus = 9
	OrderPaymentStatusCompleted OrderPaymentStatus = 10
)

type OrderPayment struct {
	ID      uint64
	OrderID uint64
	Value   float64
	Type    PaymentType
	Status  OrderPaymentStatus
}

type OrderPaymentRepository interface {
	FindPaymentByOrderAndType(ctx context.Context, orderID uint64, orderPaymentType PaymentType) (py OrderPayment, err error)
	Update(ctx context.Context, payment *OrderPayment) (err error)
	Create(ctx context.Context, payment *OrderPayment) (err error)
}
