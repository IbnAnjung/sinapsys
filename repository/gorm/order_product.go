package repository

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/orm"
	"github.com/IbnAnjung/synapsis/repository/gorm/model"
	"github.com/jinzhu/copier"
)

type orderProductRepository struct {
	uow orm.Uow
}

func NewGormOrderProductRepository(
	uow orm.Uow,
) entity.OrderProductRepository {
	return &orderProductRepository{
		uow,
	}
}

func (r *orderProductRepository) Create(ctx context.Context, order *[]entity.OrderProduct) (err error) {
	ordProds := []model.OrderProduct{}
	for _, v := range *order {
		m := model.OrderProduct{}
		m.FillFromEntity(v)
		ordProds = append(ordProds, m)
	}

	if err = r.uow.GetDB().Create(&ordProds).Error; err != nil {
		return
	}

	copier.Copy(order, &ordProds)
	return
}
