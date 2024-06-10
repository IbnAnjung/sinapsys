package router

import (
	"github.com/IbnAnjung/synapsis/cmd/http/handler"
	"github.com/IbnAnjung/synapsis/entity"

	"github.com/labstack/echo/v4"
)

func LoadAuthRouter(e *echo.Echo, authUC entity.AuthUsecase) {
	h := handler.NewAuthHandler(authUC)

	ur := e.Group("/auth")
	ur.POST("/register", h.Register)
	ur.POST("/login", h.Login)
}
