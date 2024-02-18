package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func request() {
	resp, err := http.Get("http://localhost/foo")
	if err != nil {
		log.Println("Error:", err)

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Println(string(body))
}

func main() {
	request()
}
