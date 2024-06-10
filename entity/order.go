package entity

import (
	"context"
	"time"
)

type OrderStatus uint8

const (
	OrderStatusCreated   OrderStatus = 1
	OrderStatusCanceled  OrderStatus = 9
	OrderStatusCompleted OrderStatus = 10
)

type Order struct {
	ID          uint64
	UserID      uint64
	CreatedDate time.Time
	ExpiredDate time.Time
	TotalPrice  float64
	Status      OrderStatus
}

type OrderRepository interface {
	FindByID(ctx context.Context, id uint64) (order Order, err error)
	Create(ctx context.Context, order *Order) (err error)
	Update(ctx context.Context, order *Order) (err error)
}
