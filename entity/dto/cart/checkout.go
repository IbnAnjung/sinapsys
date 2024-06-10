package cart

import "time"

type Checkout struct {
	UserID      uint64
	CartIDs     []uint64
	PaymentType uint8
}

type CheckoutOutput struct {
	OrderID     uint64
	CreatedDate time.Time
	ExpiredDate time.Time
	TotalPrice  float64
	Status      uint8
	Products    []CheckoutProduct
}

type CheckoutProduct struct {
	ProductID   uint64
	ProductName string
	Quantity    float64
	Price       float64
}
