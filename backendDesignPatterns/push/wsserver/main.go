package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var wsConnQueue = make([]*websocket.Conn, 0)

func notifyClients(lastConnData string) {
	for id, w := range wsConnQueue {
		if w == nil {
			wsConnQueue = append(wsConnQueue[:id], wsConnQueue[id+1:]...)
			continue
		}
		_, err := w.Write([]byte(lastConnData))
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func echoServer(ws *websocket.Conn) {
	remoteAddr := ws.Request().RemoteAddr
	fmt.Printf("Connected to %s\n", remoteAddr)
	wsConnQueue = append(wsConnQueue, ws)
	var message string
	if err := websocket.Message.Receive(ws, &message); err != nil {
		fmt.Println("WebSocket Read Error:", err)
		return
	}
	notifyClients("Received Data from " + remoteAddr + " : " + message)
}

func main() {
	http.Handle("/echo", websocket.Handler(echoServer))
	// Listen on :8080
	fmt.Println("Start listening TCP server on :8080 ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
