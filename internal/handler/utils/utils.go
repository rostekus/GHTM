package handler

import (
	"fmt"

	"github.com/a-h/templ"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func ValidateStruct(data interface{}) []string {
	v := validator.New()

	if err := v.Struct(data); err != nil {
		var validationErrors []string

		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}

		return validationErrors
	}

	return nil
}
