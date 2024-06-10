package cart

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/cart"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	"github.com/jinzhu/copier"
)

type CheckoutValidator struct {
	UserID      uint64   `validate:"required,numeric,min=1"`
	CartIDs     []uint64 `validate:"required,min=1"`
	PaymentType uint8    `validate:"required,min=1,oneof=1"`
}

func (uc *cartUsecase) Checkout(ctx context.Context, input cart.Checkout) (checkout cart.CheckoutOutput, err error) {
	now := uc.timeService.Now()
	validatorObj := CheckoutValidator{}
	copier.Copy(&validatorObj, &input)
	if err = uc.validator.Validate(&validatorObj); err != nil {
		return
	}

	uc.uow.Begin(ctx)
	defer uc.uow.Recovery(ctx)

	mapCarts := map[uint64]entity.Cart{}
	productIds := []uint64{}
	carts, err := uc.cartRepo.GetAndLockByIds(ctx, input.CartIDs)
	if err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	for _, v := range carts {
		mapCarts[v.ID] = v
		productIds = append(productIds, v.ProductID)
	}

	mapProducts := map[uint64]entity.Product{}
	products, err := uc.productRepo.FindByIDs(ctx, productIds)
	if err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	for _, v := range products {
		mapProducts[v.ID] = v
	}

	order := entity.Order{
		UserID:      input.UserID,
		TotalPrice:  0,
		CreatedDate: now,
		ExpiredDate: now.AddDate(0, 0, 1),
		Status:      entity.OrderStatusCreated,
	}
	orderProducts := []entity.OrderProduct{}

	for _, v := range input.CartIDs {
		exCart, ok := mapCarts[v]
		if !ok {
			uc.uow.Rollback(ctx)
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "cart product not found")
			return
		}

		if exCart.UserID != input.UserID {
			uc.uow.Rollback(ctx)
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeForbidden, "cart product not yours")
			return
		}

		exProduct, ok := mapProducts[exCart.ProductID]
		if !ok {
			uc.uow.Rollback(ctx)
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "product was deleted")
			return
		}

		order.TotalPrice += exProduct.Price * exCart.Quantity
		orderProducts = append(orderProducts, entity.OrderProduct{
			ProductID:   exProduct.ID,
			ProductName: exProduct.Name,
			Quantity:    exCart.Quantity,
			Price:       exProduct.Price,
		})
	}

	if err = uc.orderRepo.Create(ctx, &order); err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	for i := range orderProducts {
		orderProducts[i].OrderID = order.ID
	}

	if err = uc.orderProductRepo.Create(ctx, &orderProducts); err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	py := entity.OrderPayment{
		OrderID: order.ID,
		Type:    entity.PaymentType(input.PaymentType),
		Status:  entity.OrderPaymentStatusCreated,
	}

	if err = uc.orderPaymentRepo.Create(ctx, &py); err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	err = uc.cartRepo.DeleteByIds(ctx, input.CartIDs)

	uc.uow.Commit(ctx)

	checkoutProducts := []cart.CheckoutProduct{}
	for _, v := range orderProducts {
		checkoutProducts = append(checkoutProducts, cart.CheckoutProduct{
			ProductID:   v.ProductID,
			ProductName: v.ProductName,
			Quantity:    v.Quantity,
			Price:       v.Price,
		})
	}

	checkout = cart.CheckoutOutput{
		OrderID:     order.ID,
		CreatedDate: order.CreatedDate,
		ExpiredDate: order.ExpiredDate,
		TotalPrice:  order.TotalPrice,
		Status:      uint8(order.Status),
		Products:    checkoutProducts,
	}

	return
}
