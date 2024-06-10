package handler

import (
	"net/http"

	"github.com/IbnAnjung/synapsis/entity"
	paymentmanualtransfer "github.com/IbnAnjung/synapsis/entity/dto/payment_manual_transfer"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	pkgHttp "github.com/IbnAnjung/synapsis/pkg/http"

	"github.com/IbnAnjung/synapsis/cmd/http/handler/presenter"

	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	uc entity.PaymentUseCase
}

func NewPaymentHandler(
	uc entity.PaymentUseCase,
) *paymentHandler {
	return &paymentHandler{
		uc,
	}
}

func (h paymentHandler) ManualTransferConfirmation(c echo.Context) error {
	req := presenter.ManualTransferConfirmationRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	if err := h.uc.ManualTransferConfirmation(c.Request().Context(), paymentmanualtransfer.PaymentManualTransferConfirmation{
		OrderID:           req.OrderID,
		PaymentType:       req.PaymentType,
		BankAccountNumber: req.BankAccountNumber,
		BankAccountName:   req.BankAccountName,
		Date:              req.Date,
		Value:             req.Value,
	}); err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", nil))
}
