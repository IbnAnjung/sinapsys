package product

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity/dto/product"
)

type GetProductListValidator struct {
	CategoryID uint64 `validate:"numeric,omitempty"`
	Page       uint16 `validate:"numeric,omitempty"`
	Limit      uint8  `validate:"numeric,max=50,omitempty"`
}

func (uc *productUsecase) GetProductLists(ctx context.Context, input product.GetProductList) (products []product.ProductList, err error) {
	if err = uc.validator.Validate(GetProductListValidator{
		CategoryID: input.CategoryID,
		Page:       input.Page,
		Limit:      input.Limit,
	}); err != nil {
		return
	}

	if input.Limit == 0 {
		input.Limit = 10
	}

	if input.Page == 0 {
		input.Page = 1
	}

	products, err = uc.productRepo.GetProductLists(ctx, input)

	return
}
