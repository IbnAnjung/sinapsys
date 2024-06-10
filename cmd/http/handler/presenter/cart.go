package presenter

type GetCart struct {
	Product    []GetCartProduct `json:"products"`
	TotalPrice float64          `json:"total_price"`
}

type GetCartProduct struct {
	ID         uint64  `json:"cart_product_id"`
	ProductID  uint64  `json:"product_id"`
	Name       string  `json:"name"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
	TotalPrice float64 `json:"total_price"`
}

type AddCartRequest struct {
	ProductID uint64  `json:"product_id"`
	Quantity  float64 `json:"quantity"`
}

type AddCartResponse struct {
	ID        uint64  `json:"id"`
	ProductID uint64  `json:"product_id"`
	Quantity  float64 `json:"quantity"`
}

type UpdateCartRequest struct {
	ID       uint64  `param:"id"`
	Quantity float64 `json:"quantity"`
}

type UpdateCartResponse struct {
	ID        uint64  `json:"id"`
	ProductID uint64  `json:"product_id"`
	Quantity  float64 `json:"quantity"`
}

type DeleteCartRequest struct {
	ID uint64 `param:"id"`
}

type CheckoutRequest struct {
	CartsProductIds []uint64 `json:"cart_product_ids"`
	PaymentType     uint8    `json:"payment_type"`
}

type CheckoutResponse struct {
	OrderID     uint64                 `json:"order_id"`
	CreatedDate string                 `json:"created_date"`
	ExpiredDate string                 `json:"expired_date"`
	TotalPrice  float64                `json:"total_price"`
	Status      uint8                  `json:"status"`
	Products    []CheckoutOrderProduct `json:"products"`
}

type CheckoutOrderProduct struct {
	ProductID   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
}
