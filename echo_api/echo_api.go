package echo_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run_server() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}))
	v1 := e.Group("/api/v1")
	Register(v1)

	e.Logger.Fatal(e.Start(":80"))

}
