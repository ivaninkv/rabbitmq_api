package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var goroutineQty int = 16
	var messageQty int = 5000

	for i := 0; i < goroutineQty; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go DoWork(messageQty, "Text1", &wg)
	}

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")

}

func DoWork(MsgQty int, TextMessage string, wg *sync.WaitGroup) {
	defer wg.Done()

	data := url.Values{
		"name":    {"name"},
		"host":    {"localhost"},
		"port":    {"5672"},
		"user":    {"guest"},
		"pass":    {"guest"},
		"duarble": {"true"},
		"usessl":  {"false"},
	}

	resp, err := http.PostForm("http://localhost/api/v1/create_connection", data)
	if err != nil {
		fmt.Println(resp)
	}

	data = url.Values{
		"queue_id": {"1"},
		"body":     {TextMessage},
	}

	for i := 0; i < MsgQty; i++ {
		resp, err := http.PostForm("http://localhost/api/v1/send_message", data)
		if err != nil {
			log.Println("Error!")
		}

		defer resp.Body.Close()
		if _, err := io.ReadAll(resp.Body); err != nil {
			fmt.Println(err.Error())
			// } else {
			// 	fmt.Println(string(body))
		}
	}
}
