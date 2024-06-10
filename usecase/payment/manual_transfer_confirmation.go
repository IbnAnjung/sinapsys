package payment

import (
	"context"
	"time"

	"github.com/IbnAnjung/synapsis/entity"
	paymentmanualtransfer "github.com/IbnAnjung/synapsis/entity/dto/payment_manual_transfer"
	"github.com/jinzhu/copier"
)

type AddCartProductValidator struct {
	OrderID           uint64  `validate:"required"`
	PaymentType       uint8   `validate:"required,min=1,oneof=1"`
	BankAccountNumber string  `validate:"required,max=25"`
	BankAccountName   string  `validate:"required,max=50"`
	Date              string  `validate:"required,date_format=2006-01-02"`
	Value             float64 `validate:"required"`
}

func (uc *paymentUsecase) ManualTransferConfirmation(ctx context.Context, input paymentmanualtransfer.PaymentManualTransferConfirmation) (err error) {
	validatorObj := AddCartProductValidator{}
	copier.Copy(&validatorObj, &input)
	if err = uc.validator.Validate(&validatorObj); err != nil {
		return
	}

	// find Payment
	op, err := uc.orderPaymentRepo.FindPaymentByOrderAndType(ctx, input.OrderID, entity.PaymentType(input.PaymentType))
	if err != nil {
		return
	}

	d, _ := time.Parse("2006-01-02", input.Date)

	// force confirm status to completed
	py := entity.PaymentManualTransfer{
		OrderPaymentID:    op.ID,
		BankAccountNumber: input.BankAccountName,
		BankAccountName:   input.BankAccountNumber,
		Date:              d,
		Value:             input.Value,
		Status:            entity.ManualTransferPaymentStatusValid,
	}

	uc.uow.Begin(ctx)
	defer uc.uow.Recovery(ctx)

	if err = uc.paymentManualTransfer.Create(ctx, &py); err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	op.Status = entity.OrderPaymentStatusCompleted
	if err = uc.orderPaymentRepo.Update(ctx, &op); err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	order, err := uc.orderRepo.FindByID(ctx, input.OrderID)
	if err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	order.Status = entity.OrderStatusCompleted
	if err = uc.orderRepo.Update(ctx, &order); err != nil {
		uc.uow.Rollback(ctx)
		return
	}

	uc.uow.Commit(ctx)
	return
}
