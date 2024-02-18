package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"

func echo(url string) {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("C_Received: %s\n", msg[:n])

}

func chat(url string) {
	fmt.Println("Chat Client Open...")
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Channel to receiced sigint
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	// channel to stop reading input
	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				return
			case <-sigChan:
				fmt.Println("Exiting ( SIGINIT caught )")
				close(stop)
				return
			default:
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter Message: ")
				text, _ := reader.ReadString('\n')
				if text == "exit\n" {
					fmt.Println("Close received. Closing Chat")
					close(stop)
					return
				}
				if _, err := ws.Write([]byte(text)); err != nil {
					log.Println(err.Error())
					close(stop)
					return
				}

				go func() {
					var msg = make([]byte, 512)
					var n int
					if n, err = ws.Read(msg); err != nil {
						log.Fatal(err)
						close(stop)
						return
					}
					fmt.Printf("C_Received: %s\n", msg[:n])
				}()

			}
		}
	}()
	<-stop
}
func main() {
	echo("ws://localhost:12345/echo")
	chat("ws://localhost:12345/chat")
}
