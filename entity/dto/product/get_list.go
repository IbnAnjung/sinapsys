package product

type GetProductList struct {
	CategoryID uint64
	Page       uint16
	Limit      uint8
}

type ProductList struct {
	ID                uint64
	ProductCategoryID uint64
	Name              string
	Description       string
	Price             float64
	CategoryID        int64
	CategoryName      string
}
