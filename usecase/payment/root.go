package payment

import (
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/orm"
	"github.com/IbnAnjung/synapsis/pkg/structvalidator"
	"github.com/IbnAnjung/synapsis/pkg/time"
)

type paymentUsecase struct {
	timeService           time.Time
	validator             structvalidator.Validator
	uow                   orm.Uow
	orderRepo             entity.OrderRepository
	orderPaymentRepo      entity.OrderPaymentRepository
	paymentManualTransfer entity.PaymentManualTransferRepository
}

func NewUsecase(
	timeService time.Time,
	validator structvalidator.Validator,
	uow orm.Uow,
	orderRepo entity.OrderRepository,
	orderPaymentRepo entity.OrderPaymentRepository,
	paymentManualTransfer entity.PaymentManualTransferRepository,
) entity.PaymentUseCase {
	return &paymentUsecase{
		timeService,
		validator,
		uow,
		orderRepo,
		orderPaymentRepo,
		paymentManualTransfer,
	}
}
