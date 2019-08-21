package route

import (
	"database/sql"
	"net/http"

	cMiddleware "github.com/koungkub/go-structure/src/middleware"
	"github.com/koungkub/go-structure/src/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/go-playground/validator.v9"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func GetRouter(
	log *logrus.Entry,
	cache *redis.Client,
	db *sql.DB,
) *echo.Echo {

	e := echo.New()

	e.Pre(
		middleware.RemoveTrailingSlash(),
	)

	e.Use(
		middleware.Recover(),
		middleware.RequestID(),
		middleware.Secure(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{
				"*",
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
			},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentType,
				echo.HeaderAccept,
			},
			ExposeHeaders:    []string{},
			MaxAge:           3600,
			AllowCredentials: true,
		}),
	)

	e.Use(
		cMiddleware.SetConnectionMiddleware("log", log),
		cMiddleware.SetConnectionMiddleware("cache", cache),
		cMiddleware.SetConnectionMiddleware("db", db),
	)

	e.HTTPErrorHandler = utils.HTTPErrorHandler
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.GET("/_/metrics", echo.WrapHandler(promhttp.Handler()))

	e.GET("/_/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "healthy",
		})
	})

	echo := Route(e)

	return echo
}
