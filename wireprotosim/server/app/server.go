package app

import (
	"fmt"
	"log"
	"net"
)

func Main() {
	port := fmt.Sprintf("%d", 8989)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen")
	}
	defer listener.Close()
	log.Println("Server listening on port " + port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept the connection %v\n", conn)
			continue
		}
		go handleConnection(conn)
	}
}
