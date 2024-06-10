package repository

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/product"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	"github.com/IbnAnjung/synapsis/pkg/orm"
	"github.com/IbnAnjung/synapsis/repository/gorm/model"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type productRepository struct {
	uow orm.Uow
}

func NewGormProductRepository(
	uow orm.Uow,
) entity.ProductRepository {
	return &productRepository{
		uow,
	}
}

func (r *productRepository) Find(ctx context.Context, id uint64) (product entity.Product, err error) {
	m := model.Product{}
	if err = r.uow.GetDB().Where("id = ?", id).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "product not found")
		}
		return
	}

	copier.Copy(&product, &m)
	return
}

func (r *productRepository) FindByIDs(ctx context.Context, ids []uint64) (products []entity.Product, err error) {
	m := []model.Product{}
	if err = r.uow.GetDB().Where("id in (?)", ids).Find(&m).Error; err != nil {
		return
	}

	copier.Copy(&products, &m)

	return
}

func (r *productRepository) GetProductLists(ctx context.Context, input product.GetProductList) (products []product.ProductList, err error) {
	m := []model.GetProductList{}

	offset := (input.Page - 1) * uint16(input.Limit)

	q := r.uow.GetDB().Table("products p").
		Joins("JOIN product_categories pc on p.product_category_id = pc.id").
		Select(`p.*, pc.name category_name`).Offset(int(offset)).Limit(int(input.Limit))

	if input.CategoryID != 0 {
		q = q.Where("p.product_category_id = ?", input.CategoryID)
	}

	if err = q.Find(&m).Error; err != nil {
		return
	}

	for _, v := range m {
		p := product.ProductList{}
		copier.Copy(&p, &v)

		products = append(products, p)
	}

	return
}
