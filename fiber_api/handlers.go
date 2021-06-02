package fiber_api

import (
	"fmt"
	"net/http"
	"rmq_api/rmq"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetConnectionList(c *fiber.Ctx) error {
	return c.JSON(rmq.Conn)
}

func CreateConnection(c *fiber.Ctx) error {
	qp := &rmq.QueueParam{}
	if err := c.BodyParser(qp); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	conn := rmq.CreateQueue(qp)
	return c.JSON(fmt.Sprintf("{created_id:%d}", conn.Id))
}

func DeleteConnection(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	err = rmq.DeleteQueue(id)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	return c.Status(200).SendString(fmt.Sprintf("Queue %d deleted.", id))
}

func SendMessage(c *fiber.Ctx) error {
	msg := &rmq.Message{}
	if err := c.BodyParser(msg); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if err := rmq.SendMessage(msg); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	return c.Status(200).SendString("Message was be sended")
}
