package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
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
		"body":     {"my text message"},
	}

	for i := 0; i < 1000; i++ {
		resp, err := http.PostForm("http://localhost/api/v1/send_message", data)
		if err != nil {
			log.Println("Error!")
		}

		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(body))
		}
	}
}
