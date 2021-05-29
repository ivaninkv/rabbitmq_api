package echo_api

import (
	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group) {

	connList := g.Group("/connections")
	connList.GET("", GetConnectionList)

	createConn := g.Group("/create_connection")
	createConn.POST("", CreateConnection)

	deleteConn := g.Group("/delete_connection")
	deleteConn.DELETE("/:id", DeleteConnection)

	sendMsg := g.Group("/send_message")
	sendMsg.POST("", SendMessage)
}
