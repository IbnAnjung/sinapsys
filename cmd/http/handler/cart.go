package handler

import (
	"net/http"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/cart"
	pkgHttp "github.com/IbnAnjung/synapsis/pkg/http"
	"github.com/jinzhu/copier"

	"github.com/IbnAnjung/synapsis/cmd/http/handler/presenter"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	uc entity.CartUsecase
}

func NewCartHandler(
	uc entity.CartUsecase,
) *cartHandler {
	return &cartHandler{
		uc,
	}
}

func (h cartHandler) GetcartList(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)
	userID := c.Get(pkgHttp.UserIdContextKey).(int64)

	output, err := h.uc.GetCart(c.Request().Context(), uint64(userID))

	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	reProduct := []presenter.GetCartProduct{}
	copier.Copy(&reProduct, &output.Products)

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", presenter.GetCart{
		TotalPrice: output.TotalPrice,
		Product:    reProduct,
	}))
}

func (h cartHandler) Create(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	req := presenter.AddCartRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	userID := c.Get(pkgHttp.UserIdContextKey).(int64)

	output, err := h.uc.AddCartProduct(c.Request().Context(), cart.AddCartProduct{
		UserID:    uint64(userID),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	})

	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	res := presenter.AddCartResponse{}
	copier.Copy(&res, &output)

	return c.JSON(http.StatusCreated, pkgHttp.GetStandartSuccessResponse("success", res))
}

func (h cartHandler) Update(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	req := presenter.UpdateCartRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	userID := c.Get(pkgHttp.UserIdContextKey).(int64)

	output, err := h.uc.Update(c.Request().Context(), cart.UpdateCart{
		UserID:   uint64(userID),
		ID:       req.ID,
		Quantity: req.Quantity,
	})

	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	res := presenter.UpdateCartResponse{}
	copier.Copy(&res, &output)

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}

func (h cartHandler) Delete(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	req := presenter.DeleteCartRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	userID := c.Get(pkgHttp.UserIdContextKey).(int64)

	err := h.uc.Delete(c.Request().Context(), cart.DeleteCart{
		UserID: uint64(userID),
		ID:     req.ID,
	})

	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", nil))
}

func (h cartHandler) Checkout(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	req := presenter.CheckoutRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	userID := c.Get(pkgHttp.UserIdContextKey).(int64)

	output, err := h.uc.Checkout(c.Request().Context(), cart.Checkout{
		UserID:      uint64(userID),
		CartIDs:     req.CartsProductIds,
		PaymentType: req.PaymentType,
	})

	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	res := presenter.CheckoutResponse{}

	copier.Copy(&res, &output)
	res.CreatedDate = output.CreatedDate.Format("2006-01-02 15:04:05")
	res.ExpiredDate = output.ExpiredDate.Format("2006-01-02 15:04:05")

	return c.JSON(http.StatusCreated, pkgHttp.GetStandartSuccessResponse("success", res))
}
