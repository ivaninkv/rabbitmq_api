package fiber_api

import (
	"github.com/gofiber/fiber/v2"
)

func Register(g fiber.Router) {

	connList := g.Group("/connections")
	connList.Get("", GetConnectionList)

	createConn := g.Group("/create_connection")
	createConn.Post("", CreateConnection)

	deleteConn := g.Group("/delete_connection")
	deleteConn.Delete("/:id", DeleteConnection)

	sendMsg := g.Group("/send_message")
	go sendMsg.Post("", SendMessage)
}
