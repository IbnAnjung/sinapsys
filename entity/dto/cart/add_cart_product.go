package cart

type AddCartProduct struct {
	UserID    uint64
	ProductID uint64
	Quantity  float64
}

type UpdateCart struct {
	ID       uint64
	UserID   uint64
	Quantity float64
}

type DeleteCart struct {
	ID     uint64
	UserID uint64
}
