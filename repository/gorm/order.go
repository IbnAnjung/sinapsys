package repository

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/orm"
	"github.com/IbnAnjung/synapsis/repository/gorm/model"
)

type orderRepository struct {
	uow orm.Uow
}

func NewGormOrderRepository(
	uow orm.Uow,
) entity.OrderRepository {
	return &orderRepository{
		uow,
	}
}

func (r *orderRepository) FindByID(ctx context.Context, id uint64) (order entity.Order, err error) {
	m := model.Order{}

	if err = r.uow.GetDB().Where("id = ?", id).Find(&m).Error; err != nil {
		fmt.Println("=> error create order")
		return
	}

	order = m.ToEntity()

	return
}

func (r *orderRepository) Create(ctx context.Context, order *entity.Order) (err error) {
	m := model.Order{}
	m.FillFromEntity(*order)

	if err = r.uow.GetDB().Create(&m).Error; err != nil {
		return
	}

	order.ID = m.ID

	return
}

func (r *orderRepository) Update(ctx context.Context, order *entity.Order) (err error) {
	m := model.Order{}
	m.FillFromEntity(*order)

	if err = r.uow.GetDB().Updates(&m).Error; err != nil {
		return
	}

	return
}
