package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rostekus/ghtm/internal/app/service/dto"
	"github.com/rostekus/ghtm/internal/app/service/ports/output"
	handler "github.com/rostekus/ghtm/internal/handler/utils"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	Service output.App
}

func (h *Handler) RegisterUserApi(c echo.Context) error {
	var userRegisterRequest UserRegisterRequest
	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&userRegisterRequest)
	if err != nil {
		log.Error().Err(err).Msg("handler:RegisterUserApi:JsonDecoder")
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "cannot parse the body"})
		return err
	}
	errs := handler.ValidateStruct(userRegisterRequest)
	if errs != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: strings.Join(errs, ",")})
		return fmt.Errorf("bad request")
	}
	user := dto.UserDTO{
		Email:    userRegisterRequest.Email,
		Password: userRegisterRequest.Password,
		Username: userRegisterRequest.Username,
	}
	u, err := h.Service.CreateUser(c.Request().Context(), user)
	if err != nil {
		log.Error().Err(err).Msg("handler:RegisterUserApi:CreateUser")
		c.JSON(http.StatusInternalServerError, handler.ErrorResponse{Message: "cannot create user"})
		return err
	}
	c.JSON(http.StatusOK, UserRegisterResponse{
		User: u,
	})
	return nil
}

func (h *Handler) LoginUserApi(c echo.Context) error {
	var userLoginRequest UserLoginRequest

	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(&userLoginRequest)
	if err != nil {
		log.Error().Err(err).Msg("handler:RegisterUserApi:JsonDecoder")
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: "cannot parse the body"})
		return err
	}

	errs := handler.ValidateStruct(userLoginRequest)
	if errs != nil {
		log.Error().Err(err).Msg("handler:RegisterUserApi:ValudateStruct")
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{Message: strings.Join(errs, ",")})
		return fmt.Errorf("bad request")
	}

	user := dto.UserDTO{
		Email:    userLoginRequest.Email,
		Password: userLoginRequest.Password,
	}
	user, err = h.Service.GetUser(c.Request().Context(), user)

	if err != nil {
		http.Error(c.Response(), err.Error(), http.StatusUnauthorized)
		return err
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		User: user,
	})

	return nil
}
