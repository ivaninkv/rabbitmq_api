package rmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CreateQueue(qp *QueueParam) Connections {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", qp.User, qp.Pass, qp.Host, qp.Port))
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	_, err = ch.QueueDeclare(
		qp.Name,    // name
		qp.Durable, // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	c := Connections{getMaxQueueNum() + 1, qp.Name, qp.Host, qp.Durable, conn, ch}
	Conn = append(Conn, c)

	return c
}

func DeleteQueue(Id int) (err error) {
	indexToDel := getIndexByQueueNum(Id)
	if indexToDel >= 0 {
		ch := Conn[indexToDel].Channel
		ch.Close()
		conn := Conn[indexToDel].Conn
		conn.Close()
		Conn = append(Conn[:indexToDel], Conn[indexToDel+1:]...)
	} else {
		err = fmt.Errorf("queue %d not found", Id)
		log.Printf("Index %d not found", Id)
	}

	return err
}

func SendMessage(msg *Message) (err error) {
	idx := getIndexByQueueNum(msg.Id)
	if idx < 0 {
		failOnError(fmt.Errorf("index less then 0"), "Failed to get a channel")
	}

	ch, err := Conn[idx].Conn.Channel()
	failOnError(err, "Failed to get a channel")

	err = ch.Publish(
		"",             // exchange
		Conn[idx].Name, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg.Body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", msg.Body)

	return err
}

func getMaxQueueNum() int {
	maxId := 0
	for _, v := range Conn {
		if v.Id > maxId {
			maxId = v.Id
		}
	}

	return maxId
}

func getIndexByQueueNum(Id int) int {
	var indexToDel int = -1
	for k, v := range Conn {
		if v.Id == Id {
			indexToDel = k
		}
	}

	return indexToDel
}
