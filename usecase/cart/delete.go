package cart

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity/dto/cart"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	"github.com/jinzhu/copier"
)

type DeleteCartValidator struct {
	ID     uint64 `validate:"required,numeric,min=1"`
	UserID uint64 `validate:"required,numeric,min=1"`
}

func (uc *cartUsecase) Delete(ctx context.Context, input cart.DeleteCart) (err error) {
	validatorObj := DeleteCartValidator{}
	copier.Copy(&validatorObj, &input)
	if err = uc.validator.Validate(&validatorObj); err != nil {
		return
	}

	uc.uow.Begin(ctx)
	defer uc.uow.Recovery(ctx)

	cart, err := uc.cartRepo.FindAndLock(ctx, input.ID)
	if err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	if cart.UserID != input.UserID {
		uc.uow.Rollback(ctx)
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeForbidden, "access denied")
		return
	}

	if err = uc.cartRepo.Delete(ctx, input.ID); err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	uc.uow.Commit(ctx)

	return
}
