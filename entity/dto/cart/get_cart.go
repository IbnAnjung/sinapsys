package cart

type GetCart struct {
	Products   []GetCartProduct
	TotalPrice float64
}

type GetCartProduct struct {
	ID         uint64
	ProductID  uint64
	Name       string
	Quantity   string
	Price      float64
	TotalPrice float64
}
