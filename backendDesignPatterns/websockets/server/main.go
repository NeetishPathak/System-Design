package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Println("Error reading the message", err)
		return
	}
	fmt.Println("S_Received:", string(msg[:n]))

	if _, err = ws.Write(msg); err != nil {
		log.Println("Error sending the messgae", err)
		return
	}
}

func ChatServer(ws *websocket.Conn) {
	defer ws.Close()

	for {
		msg := make([]byte, 512)
		n, err := ws.Read(msg)
		if err != nil {
			log.Println("Error reading the message", err)
			return
		}
		fmt.Println("S_Received:", string(msg[:n]))

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Server Message: ")
		text, _ := reader.ReadString('\n')
		if _, err := ws.Write([]byte(text)); err != nil {
			log.Println(err.Error())
			return
		}
	}

}

func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	http.Handle("/chat", websocket.Handler(ChatServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal(err)
	}
}
