package repository

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/cart"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	"github.com/IbnAnjung/synapsis/pkg/orm"
	"github.com/IbnAnjung/synapsis/repository/gorm/model"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"gorm.io/gorm/clause"
)

type cartRepository struct {
	uow orm.Uow
}

func NewGormCartRepository(
	uow orm.Uow,
) entity.CartRepository {
	return &cartRepository{
		uow,
	}
}

func (r *cartRepository) AddProduct(ctx context.Context, input *entity.Cart) (err error) {
	m := model.Cart{}
	copier.Copy(&m, &input)

	err = r.uow.GetDB().Select("quantity", "user_id", "product_id").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "product_id"}, {Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"quantity": r.uow.GetDB().Raw(fmt.Sprintf("quantity+%02f", input.Quantity))}),
	}).
		Create(&m).Error

	input.ID = m.ID
	input.Quantity = m.Quantity

	return
}

func (r *cartRepository) Update(ctx context.Context, input entity.Cart) (err error) {
	m := model.Cart{}
	copier.Copy(&m, &input)

	err = r.uow.GetDB().Updates(&m).Error

	return
}

func (r *cartRepository) Delete(ctx context.Context, id uint64) (err error) {
	m := model.Cart{ID: uint64(id)}

	return r.uow.GetDB().Delete(&m).Error
}

func (r *cartRepository) DeleteByIds(ctx context.Context, ids []uint64) (err error) {
	return r.uow.GetDB().Where("id in (?)", ids).Delete(&model.Cart{}).Error
}

func (r *cartRepository) GetAndLockByIds(ctx context.Context, ids []uint64) (carts []entity.Cart, err error) {
	m := []model.Cart{}
	if err = r.uow.GetDB().Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id IN (?)", ids).Find(&m).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "cart product not found")
		}
		return
	}

	copier.Copy(&carts, &m)

	return
}

func (r *cartRepository) FindAndLock(ctx context.Context, id uint64) (cart entity.Cart, err error) {
	m := model.Cart{}
	if err = r.uow.GetDB().Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", id).First(&m).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "cart product not found")
		}
		return
	}

	copier.Copy(&cart, &m)
	return
}

func (r *cartRepository) Find(ctx context.Context, id uint64) (cart entity.Cart, err error) {
	m := model.Cart{}

	if err = r.uow.GetDB().Where("id = ?", id).First(&m).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "cart not found")
		}
		return
	}

	copier.Copy(&cart, &m)
	return
}

func (r *cartRepository) GetCartProducts(ctx context.Context, userID uint64) (cartProduct []cart.GetCartProduct, err error) {
	m := []model.GetCart{}

	if err = r.uow.GetDB().Table("carts c").
		Joins("JOIN products p ON c.product_id = p.id").
		Select(`c.id,
			p.id product_id,
			p.name,
			c.quantity,
			p.price,
			p.price*c.quantity total_price`).
		Where("c.user_id = ?", userID).Find(&m).Error; err != nil {
		return
	}

	for _, v := range m {
		c := cart.GetCartProduct{}
		copier.Copy(&c, &v)
		cartProduct = append(cartProduct, c)
	}

	return
}
