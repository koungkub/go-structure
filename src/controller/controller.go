package controller

import (
	"net/http"

	"github.com/koungkub/go-structure/src/service"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func GetNameController(user service.User) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := user.GetName()
		return c.JSON(http.StatusOK, echo.Map{
			"message": name,
		})
	}
}

func SetNameController(user service.User, validate *validator.Validate) echo.HandlerFunc {
	return func(c echo.Context) error {

		name := c.Param("name")
		err := validate.Var(name, "alphanum")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "name not a alpha and num")
		}

		user.SetName(name)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "can not set name")
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"name": user.GetName(),
		})
	}
}
