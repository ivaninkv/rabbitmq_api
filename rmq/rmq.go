package rmq

import "github.com/streadway/amqp"

type Connections struct {
	Id      int              `json:"id" xml:"id" form:"id" query:"id"`
	Name    string           `json:"name" xml:"name" form:"name" query:"name"`
	Host    string           `json:"host" xml:"host" form:"host" query:"host"`
	Durable bool             `json:"durable" xml:"durable" form:"durable" query:"durable"`
	Conn    *amqp.Connection `json:"-" xml:"-" form:"-" query:"-"`
	Channel *amqp.Channel    `json:"-" xml:"-" form:"-" query:"-"`
}

type QueueParam struct {
	Name    string `json:"name" xml:"name" form:"name" query:"name"`
	Host    string `json:"host" xml:"host" form:"host" query:"host"`
	Port    int    `json:"port" xml:"port" form:"port" query:"port"`
	User    string `json:"user" xml:"user" form:"user" query:"user"`
	Pass    string `json:"pass" xml:"pass" form:"pass" query:"pass"`
	Durable bool   `json:"durable" xml:"duarble" form:"duarble" query:"duarble"`
	UseSsl  bool   `json:"usessl" xml:"usessl" form:"usessl" query:"usessl"`
}

type Message struct {
	QueueId int    `json:"queue_id" xml:"queue_id" form:"queue_id" query:"queue_id"`
	Body    string `json:"body" xml:"body" form:"body" query:"body"`
}

var Conn = make([]Connections, 0)
