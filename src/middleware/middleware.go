package middleware

import (
	"github.com/labstack/echo/v4"
)

func SetConnectionMiddleware(name string, middleware interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			c.Set(name, middleware)
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		})
	}
}
