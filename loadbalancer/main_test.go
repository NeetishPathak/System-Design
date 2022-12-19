package main

import (
	"fmt"
	"log"
	"net"
	"testing"
)

type Client struct {
	id  int
	msg string
}

func TestDial(t *testing.T) {
	clients := []Client{
		{1, "Android"},
		{2, "Iphone"},
		{3, "Windows"},
		{4, "MacOS"},
		{5, "Ubuntu"},
		{6, "SUSE"},
		{7, "Kali"},
		{8, "Fedora"},
	}
	// Dial a connection
	for _, client := range clients {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		fmt.Println("Client Connected Successfully", client)
		text := []byte("Hello " + client.msg)
		_, err = conn.Write(text)
		if err != nil {
			log.Panicln("Client write error", err)
		}
		fmt.Println("Snt: ", string(text))

		//read response
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Client Read error", err)
		}

		fmt.Println("Rcv: ", string(buf[:n]))
	}
}
