package handler

import (
	"net/http"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/product"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	pkgHttp "github.com/IbnAnjung/synapsis/pkg/http"
	"github.com/jinzhu/copier"

	"github.com/IbnAnjung/synapsis/cmd/http/handler/presenter"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	uc entity.ProductUsecase
}

func NewProductHandler(
	uc entity.ProductUsecase,
) *productHandler {
	return &productHandler{
		uc,
	}
}
func (h productHandler) GetProductList(c echo.Context) error {
	req := presenter.GetProductListRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	output, err := h.uc.GetProductLists(c.Request().Context(), product.GetProductList{
		CategoryID: uint64(req.CategoryID),
		Page:       uint16(req.Page),
		Limit:      uint8(req.Limit),
	})

	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	response := []presenter.GetProductListResponse{}
	for _, v := range output {
		resProd := presenter.GetProductListResponse{}
		copier.Copy(&resProd, &v)

		response = append(response, resProd)
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", response))
}
