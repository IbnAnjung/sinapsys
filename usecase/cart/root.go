package cart

import (
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/orm"
	"github.com/IbnAnjung/synapsis/pkg/structvalidator"
	"github.com/IbnAnjung/synapsis/pkg/time"
)

type cartUsecase struct {
	timeService      time.Time
	validator        structvalidator.Validator
	cartRepo         entity.CartRepository
	uow              orm.Uow
	productRepo      entity.ProductRepository
	orderRepo        entity.OrderRepository
	orderProductRepo entity.OrderProductRepository
	orderPaymentRepo entity.OrderPaymentRepository
}

func NewUsecase(
	timeService time.Time,
	validator structvalidator.Validator,
	cartRepo entity.CartRepository,
	uow orm.Uow,
	productRepo entity.ProductRepository,
	orderRepo entity.OrderRepository,
	orderProductRepo entity.OrderProductRepository,
	orderPaymentRepo entity.OrderPaymentRepository,
) entity.CartUsecase {
	return &cartUsecase{
		timeService,
		validator,
		cartRepo,
		uow,
		productRepo,
		orderRepo,
		orderProductRepo,
		orderPaymentRepo,
	}
}
