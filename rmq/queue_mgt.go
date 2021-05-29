package rmq

import (
	"fmt"
	"log"
)

func CreateQueue(qp *QueueParam) Connections {
	c := Connections{getMaxQueueNum() + 1, qp.Name, qp.Host, qp.Durable}
	Conn = append(Conn, c)

	return c
}

func DeleteQueue(Id int) (err error) {
	indexToDel := getIndexByQueueNum(Id)
	if indexToDel >= 0 {
		Conn = append(Conn[:indexToDel], Conn[indexToDel+1:]...)
	} else {
		err = fmt.Errorf("queue %d not found", Id)
		log.Printf("Index %d not found", Id)
	}

	return err
}

func SendMessage(msg *Message) (err error) {

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
