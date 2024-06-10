package cart

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity/dto/cart"
)

func (uc *cartUsecase) GetCart(ctx context.Context, userID uint64) (cart cart.GetCart, err error) {

	cartProducts, err := uc.cartRepo.GetCartProducts(ctx, userID)
	if err != nil {
		return cart, err
	}

	for _, v := range cartProducts {
		cart.TotalPrice += v.TotalPrice
	}

	cart.Products = cartProducts

	return
}
