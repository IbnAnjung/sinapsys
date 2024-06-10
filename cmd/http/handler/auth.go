package handler

import (
	"net/http"

	"github.com/IbnAnjung/synapsis/entity"
	"github.com/IbnAnjung/synapsis/entity/dto/auth"
	coreerror "github.com/IbnAnjung/synapsis/pkg/error"
	pkgHttp "github.com/IbnAnjung/synapsis/pkg/http"

	"github.com/IbnAnjung/synapsis/cmd/http/handler/presenter"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	uc entity.AuthUsecase
}

func NewAuthHandler(
	uc entity.AuthUsecase,
) *authHandler {
	return &authHandler{
		uc,
	}
}

func (h authHandler) Register(c echo.Context) error {
	req := presenter.RegisterRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		return err
	}

	output, err := h.uc.RegisterUser(c.Request().Context(), auth.RegisterInput{
		Name:            req.Name,
		PhoneNumber:     req.PhoneNumber,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	res := presenter.RegisterResponse{
		ID:           output.ID,
		Name:         output.Name,
		PhoneNumber:  output.PhoneNumber,
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	}

	return c.JSON(http.StatusCreated, pkgHttp.GetStandartSuccessResponse("success", res))
}

func (h authHandler) Login(c echo.Context) error {
	req := presenter.LoginRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	output, err := h.uc.Login(c.Request().Context(), auth.LoginInput{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	res := presenter.LoginResponse{
		ID:           output.ID,
		Name:         output.Name,
		PhoneNumber:  output.PhoneNumber,
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}
