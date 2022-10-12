package router

import (
	"github.com/labstack/echo"
)

func NewEchoRouter() *echo.Echo {
	e := echo.New()
	for _, route := range routes {
		switch route.Method {
		case "GET":
			e.GET(route.Path, route.EchoHandlerFunc)
		}
	}
	return e
}
