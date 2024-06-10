package entity

import "context"

type OrderProduct struct {
	ID          uint64
	OrderID     uint64
	ProductID   uint64
	ProductName string
	Quantity    float64
	Price       float64
}

type OrderProductRepository interface {
	Create(ctx context.Context, products *[]OrderProduct) (err error)
}
