package main

import (
	"fmt"
	"log"
	"net"

	"github.com/NeetishPathak/wireprotosim/server/proto"
)

func main() {
	port := fmt.Sprintf("%d", 8989)
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		log.Println("Failed to connect to the server")
		return
	}
	defer conn.Close()

	// Create a ping request pkt
	sendRequest(&conn, proto.ReqPacket{ReqType: proto.PingReq, Data: []int32{0, 0}, Message: []byte("Ping")})

	// Read ping response
	readResponse(&conn)

	// Create a add request pkt
	sendRequest(&conn, proto.ReqPacket{ReqType: proto.AddReq, Data: []int32{2, 6}, Message: []byte("Add")})

	// Read add response
	readResponse(&conn)
}

func readResponse(conn *net.Conn) {
	var respPkt proto.RespPacket
	if err := proto.ReadResponse(conn, &respPkt); err != nil {
		log.Println("Error: ", err)
	}
	if respPkt.Status == 0 {
		log.Println("Client: Received " + string(respPkt.Message))
		if respPkt.RespType == proto.PingReq {
		} else if respPkt.RespType == proto.AddReq {
			log.Println("Client: Result " + fmt.Sprintf("%d", respPkt.Result))
		}
	}
}

func sendRequest(conn *net.Conn, reqPkt proto.ReqPacket) {
	if err := proto.WriteRequest(conn, reqPkt); err != nil {
		log.Println(err)
	} else {
		log.Println("Client: Send", string(reqPkt.Message))
	}

}
