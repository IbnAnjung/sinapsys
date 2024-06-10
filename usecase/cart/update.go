package cart

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/cart"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	"github.com/jinzhu/copier"
)

type UpdateCartValidator struct {
	ID       uint64  `validate:"required,numeric,min=1"`
	UserID   uint64  `validate:"required,numeric,min=1"`
	Quantity float64 `validate:"required,numeric,min=1"`
}

func (uc *cartUsecase) Update(ctx context.Context, input cart.UpdateCart) (cart entity.Cart, err error) {
	validatorObj := UpdateCartValidator{}
	copier.Copy(&validatorObj, &input)
	if err = uc.validator.Validate(&validatorObj); err != nil {
		return
	}

	uc.uow.Begin(ctx)
	defer uc.uow.Recovery(ctx)

	cart, err = uc.cartRepo.FindAndLock(ctx, input.ID)
	if err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	if cart.UserID != input.UserID {
		uc.uow.Rollback(ctx)
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeForbidden, "access denied")
		return
	}

	cart.Quantity += input.Quantity

	if err = uc.cartRepo.Update(ctx, cart); err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	uc.uow.Commit(ctx)

	return
}
