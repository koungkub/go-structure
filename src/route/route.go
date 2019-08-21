package route

import (
	"github.com/koungkub/go-structure/src/controller"
	"github.com/koungkub/go-structure/src/service"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func Route(e *echo.Echo) *echo.Echo {

	api := e.Group("/api/v1")
	{
		api.GET("/name", controller.GetNameController(new(service.Admin)))
		api.GET("/name/:name", controller.SetNameController(new(service.Admin), validator.New()))
	}

	return e
}
