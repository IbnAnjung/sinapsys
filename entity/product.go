package entity

import (
	"context"

	"github.com/IbnAnjung/synapsis/entity/dto/product"
)

type Product struct {
	ID                uint64
	ProductCategoryID uint64
	Name              string
	Description       string
	Price             float64
}

type ProductUsecase interface {
	GetProductLists(ctx context.Context, input product.GetProductList) (products []product.ProductList, err error)
}

type ProductRepository interface {
	Find(ctx context.Context, id uint64) (product Product, err error)
	FindByIDs(ctx context.Context, id []uint64) (products []Product, err error)
	GetProductLists(ctx context.Context, input product.GetProductList) (products []product.ProductList, err error)
}
