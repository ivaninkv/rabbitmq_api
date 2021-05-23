package echo_api

import (
	"fmt"
	"net/http"
	"rmq_api/rmq"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetConnectionList(c echo.Context) error {
	return c.JSON(http.StatusOK, rmq.Conn)
}

func CreateConnection(c echo.Context) error {
	qp := &rmq.QueueParam{}
	if err := c.Bind(qp); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	conn := rmq.CreateQueue(qp)
	return c.JSON(http.StatusOK, conn)
}

func DeleteConnection(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = rmq.DeleteQueue(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("Queue %d deleted.", id))
}
