package entity

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity/dto/cart"
)

type Cart struct {
	ID        uint64
	UserID    uint64
	ProductID uint64
	Quantity  float64
}

type CartUsecase interface {
	GetCart(ctx context.Context, userID uint64) (cartProduct cart.GetCart, err error)
	AddCartProduct(ctx context.Context, input cart.AddCartProduct) (cart Cart, err error)
	Update(ctx context.Context, input cart.UpdateCart) (cart Cart, err error)
	Delete(ctx context.Context, input cart.DeleteCart) (err error)
	Checkout(ctx context.Context, input cart.Checkout) (order cart.CheckoutOutput, err error)
}

type CartRepository interface {
	GetCartProducts(ctx context.Context, userID uint64) (cartProduct []cart.GetCartProduct, err error)
	AddProduct(ctx context.Context, input *Cart) (err error)
	Update(ctx context.Context, cart Cart) (err error)
	Delete(ctx context.Context, id uint64) (err error)
	FindAndLock(ctx context.Context, id uint64) (cart Cart, err error)
	GetAndLockByIds(ctx context.Context, ids []uint64) (carts []Cart, err error)
	DeleteByIds(ctx context.Context, ids []uint64) (err error)
	Find(ctx context.Context, id uint64) (cart Cart, err error)
}
