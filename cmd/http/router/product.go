package router

import (
	"github.com/IbnAnjung/synapsis/cmd/http/handler"
	"github.com/IbnAnjung/synapsis/entity"

	"github.com/labstack/echo/v4"
)

func LoadProductRouter(e *echo.Echo, uc entity.ProductUsecase) {
	h := handler.NewProductHandler(uc)

	ur := e.Group("/product")
	ur.GET("", h.GetProductList)
}
