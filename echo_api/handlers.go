package echo_api

import (
	"net/http"
	"rmq_api/rmq"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "text")
}

func GetConnectionList(c echo.Context) error {
	return c.JSON(http.StatusOK, rmq.Conn)
}
