package cart

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/cart"
	"github.com/jinzhu/copier"
)

type AddCartProductValidator struct {
	UserID    uint64  `validate:"required,numeric,min=1"`
	ProductID uint64  `validate:"required,numeric,min=1"`
	Quantity  float64 `validate:"required,numeric,min=1"`
}

func (uc *cartUsecase) AddCartProduct(ctx context.Context, input cart.AddCartProduct) (cart entity.Cart, err error) {
	validatorObj := AddCartProductValidator{}
	copier.Copy(&validatorObj, &input)
	if err = uc.validator.Validate(&validatorObj); err != nil {
		return
	}

	// validate product
	if _, err = uc.productRepo.Find(ctx, input.ProductID); err != nil {
		return
	}

	copier.Copy(&cart, &input)

	if err = uc.cartRepo.AddProduct(ctx, &cart); err != nil {
		return
	}

	cart, err = uc.cartRepo.Find(ctx, cart.ID)

	return
}
