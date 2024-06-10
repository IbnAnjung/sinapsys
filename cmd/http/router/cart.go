package router

import (
	"github.com/IbnAnjung/synapsis/cmd/http/handler"
	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/pkg/http/echomiddleware"
	"github.com/IbnAnjung/synapsis/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func LoadCartRouter(e *echo.Echo, jwtService jwt.JwtService, uc entity.CartUsecase) {
	h := handler.NewCartHandler(uc)

	authMidWare := echomiddleware.AuthenticationMiddleware(jwtService)

	ur := e.Group("/cart", authMidWare)
	ur.POST("/checkout", h.Checkout)
	ur.GET("", h.GetcartList)
	ur.POST("", h.Create)
	ur.PUT("/:id", h.Update)
	ur.DELETE("/:id", h.Delete)
}
