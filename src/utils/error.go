package utils

import (
	"github.com/labstack/echo/v4"
)

func HTTPErrorHandler(err error, c echo.Context) {

	if status, ok := err.(*echo.HTTPError); ok {
		c.JSON(status.Code, echo.Map{
			"message": status.Message,
		})
	}
}
