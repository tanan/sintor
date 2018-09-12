package web

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tanan/sintor/config"
)

func LoadRouter(con *config.Config) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/system/health", ping)

	return e
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}