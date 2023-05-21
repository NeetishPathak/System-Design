package app

import (
	"log"
	"net"

	"github.com/NeetishPathak/wireprotosim/server/proto"
)

// To handle the connection
// server will first read the requestCode
// based on the request code type, it will call the appropriate handler

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {

		var reqPkt proto.ReqPacket
		respPacket := proto.RespPacket{Result: 0, Status: -1, Message: []byte("")}

		if err := proto.ReadRequest(&conn, &reqPkt); err != nil {
			return
		}

		reqCode := reqPkt.ReqType
		data := reqPkt.Data
		message := reqPkt.Message
		log.Println("Reading client request")
		switch reqCode {
		case proto.PingReq:
			respPacket.RespType = proto.PingReq
			respPacket.Message = handlePing(message)
			respPacket.Status = 0
		case proto.AddReq:
			respPacket.RespType = proto.AddReq
			respPacket.Result = handleAdd(data, message)
			respPacket.Message = []byte("Add-Result")
			respPacket.Status = 0
		default:
			log.Println("Invalid Request")
		}

		if err := proto.WriteResponse(&conn, respPacket); err != nil {
			log.Printf("Error writing the response: %s", err.Error())
			return
		}
	}

}

func handlePing(message []byte) []byte {
	log.Println("Server: Received " + string(message))
	sendMsg := "Pong"
	log.Println("Server: Send " + sendMsg)
	return []byte(sendMsg)
}

func handleAdd(data []int32, message []byte) int32 {
	log.Println("Server: Received " + string(message))
	var sum int32 = -1
	for _, i := range data {
		sum += i
	}
	return sum
}
