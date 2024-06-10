package repository

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/orm"
	"github.com/IbnAnjung/synapsis/repository/gorm/model"
)

type paymentManualTransferRepository struct {
	uow orm.Uow
}

func NewGormPaymentManualTransferRepository(
	uow orm.Uow,
) entity.PaymentManualTransferRepository {
	return &paymentManualTransferRepository{
		uow,
	}
}

func (r *paymentManualTransferRepository) Create(ctx context.Context, py *entity.PaymentManualTransfer) (err error) {
	m := model.PaymentManualTransfer{}
	m.FillFromEntity(*py)

	if err = r.uow.GetDB().Create(&m).Error; err != nil {
		return
	}

	py.ID = m.ID

	return
}
