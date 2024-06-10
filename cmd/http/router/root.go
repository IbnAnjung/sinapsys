package router

import (
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func SetupRouter(
	e *echo.Echo,
	jwtService jwt.JwtService,
	authUc entity.AuthUsecase,
	prodUc entity.ProductUsecase,
	cartUc entity.CartUsecase,
	paymentUc entity.PaymentUseCase,
) {

	LoadHealtRouter(e)
	LoadAuthRouter(e, authUc)
	LoadProductRouter(e, prodUc)
	LoadCartRouter(e, jwtService, cartUc)
	LoadPaymentRouter(e, jwtService, paymentUc)
}
