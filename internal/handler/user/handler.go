package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rostekus/ghtm/internal/app/service/dto"
	"github.com/rostekus/ghtm/internal/app/service/ports/output"
	handler "github.com/rostekus/ghtm/internal/handler/utils"
)

type Handler struct {
	Service output.App
}

func (h *Handler) RegisterUserApi(c echo.Context) error {
	var userRegisterRequest UserRegisterRequest

	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&userRegisterRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "cannot parse the body"})
		return err
	}

	return nil
}

func (h *Handler) LoginUserApi(c echo.Context) error {
	var userRegisterRequest UserLoginRequest

	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&userRegisterRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "cannot parse the body"})
		return err
	}
	user := dto.UserDTO{
		Email:    userRegisterRequest.Email,
		Password: userRegisterRequest.Password,
	}

	user, err = h.Service.GetUser(user)

	c.JSON(http.StatusOK, UserLoginResponse{
		User: user,
	})

	return nil
}
