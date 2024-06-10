package product

import (
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/structvalidator"
)

type productUsecase struct {
	validator   structvalidator.Validator
	productRepo entity.ProductRepository
}

func NewUsecase(
	validator structvalidator.Validator,
	productRepo entity.ProductRepository,
) entity.ProductUsecase {
	return &productUsecase{
		validator,
		productRepo,
	}
}
