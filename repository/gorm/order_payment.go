package repository

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/orm"
	"github.com/IbnAnjung/synapsis/repository/gorm/model"
)

type orderPaymentRepository struct {
	uow orm.Uow
}

func NewGormOrderPaymentRepository(
	uow orm.Uow,
) entity.OrderPaymentRepository {
	return &orderPaymentRepository{
		uow,
	}
}

func (r *orderPaymentRepository) FindPaymentByOrderAndType(ctx context.Context, orderID uint64, orderPaymentType entity.PaymentType) (orderPy entity.OrderPayment, err error) {
	m := model.OrderPayment{}

	if err = r.uow.GetDB().Where("order_id = ?", orderID).
		Where("type = ?", orderPaymentType).
		Find(&m).Error; err != nil {
		return
	}

	orderPy = m.ToEntity()

	return
}

func (r *orderPaymentRepository) Create(ctx context.Context, orderPy *entity.OrderPayment) (err error) {
	m := model.OrderPayment{}
	m.FillFromEntity(*orderPy)

	if err = r.uow.GetDB().Create(&m).Error; err != nil {
		return
	}

	orderPy.ID = m.ID

	return
}

func (r *orderPaymentRepository) Update(ctx context.Context, orderPy *entity.OrderPayment) (err error) {
	m := model.OrderPayment{}
	m.FillFromEntity(*orderPy)

	if err = r.uow.GetDB().Updates(&m).Error; err != nil {
		return
	}

	orderPy.ID = m.ID

	return
}
