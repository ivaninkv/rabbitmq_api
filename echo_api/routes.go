package echo_api

import (
	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group) {
	h := g.Group("/hello")
	h.GET("", hello)

	ConnectionLst := g.Group("/connections")
	ConnectionLst.GET("", GetConnectionList)
}
