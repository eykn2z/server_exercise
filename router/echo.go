package router

import (
	"server/handler"

	"github.com/labstack/echo"
)

func NewEchoRouter() *echo.Echo {
	e := echo.New()

	e.GET("/persons", handler.GetEchoPersons)

	return e
}
