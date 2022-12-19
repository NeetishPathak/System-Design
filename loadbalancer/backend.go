package main

import (
	"fmt"
	"log"
	"net"
)

type Backend struct {
	ip      string
	port    string
	name    string
	enabled bool
	ratio   int
}

func (bk Backend) printBackend() {
	fmt.Printf("Backed Info - Name: %s, Endpoint: %s:%s, Enabled:%v, Ratio: %d\n", bk.name, bk.ip, bk.port, bk.enabled, bk.ratio)
}

func (bk Backend) handleConn(conn net.Conn) {
	defer conn.Close()
	// read bytes
	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		log.Println(bk.name, " - failed to read: ", err)
		return
	}
	responseStr := fmt.Sprintf("%v-%v\n", bk.name, string(buf[:n]))
	_, err = conn.Write([]byte(responseStr))
	if err != nil {
		log.Println()
	}
}

func (bk Backend) Initialize() {
	laddr, err := net.ResolveTCPAddr("tcp", bk.ip+":"+bk.port)
	if err != nil {
		log.Panicln(err)
	}
	listener, err := net.Listen("tcp", laddr.String())
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()

	//req/response loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("failed to accept conn:", err)
			conn.Close()
			continue
		}
		log.Println(bk.name, " connected to: ", conn.RemoteAddr())
		go bk.handleConn(conn)
	}
}
