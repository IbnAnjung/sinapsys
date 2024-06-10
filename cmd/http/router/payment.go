package router

import (
	"github.com/IbnAnjung/synapsis/cmd/http/handler"
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/http/echomiddleware"
	"github.com/IbnAnjung/synapsis/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func LoadPaymentRouter(e *echo.Echo, jwtService jwt.JwtService, uc entity.PaymentUseCase) {
	h := handler.NewPaymentHandler(uc)

	authMidWare := echomiddleware.AuthenticationMiddleware(jwtService)

	ur := e.Group("/payment", authMidWare)
	ur.POST("/manual-transfer/confirm", h.ManualTransferConfirmation)
}
