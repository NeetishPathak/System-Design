package proto

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

const (
	PingReq = 1
	AddReq  = 2
)

type ReqPacket struct {
	ReqType int32
	Data    []int32
	Message []byte
}

type RespPacket struct {
	RespType int32
	Result   int32
	Status   int32
	Message  []byte
}

func ReadRequest(conn *net.Conn, reqPkt *ReqPacket) error {
	buf := make([]byte, 1024)
	n, err := (*conn).Read(buf)
	if err != nil {
		if err == io.EOF {
			log.Println("Connection closed by client")
			return err
		}
		log.Println("Error reading the request:", err)
		return err
	}
	dataBuf := bytes.NewReader(buf[:n])
	if err := binary.Read(dataBuf, binary.BigEndian, &reqPkt.ReqType); err != nil {
		log.Println("Error Reading the request", err)
		return err
	}

	var dataLength int32
	if err := binary.Read(dataBuf, binary.BigEndian, &dataLength); err != nil {
		log.Println("Error reading Data length:", err)
		return err
	}
	reqPkt.Data = make([]int32, dataLength)
	for i := 0; i < int(dataLength); i++ {
		if err := binary.Read(dataBuf, binary.BigEndian, &reqPkt.Data[i]); err != nil {
			log.Println("Error reading Data element:", err)
			return err
		}
	}
	messageBytes := make([]byte, dataBuf.Len())
	if err := binary.Read(dataBuf, binary.BigEndian, messageBytes); err != nil {
		log.Println("Error reading Message:", err)
		return err
	}
	reqPkt.Message = []byte(messageBytes)
	return nil
}

func WriteRequest(conn *net.Conn, reqPkt ReqPacket) error {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, reqPkt.ReqType); err != nil {
		log.Println("Error Writing the reqPkt.ReqType", err)
		return err
	}
	var dataLength int32 = int32(len(reqPkt.Data))
	if err := binary.Write(buf, binary.BigEndian, dataLength); err != nil {
		log.Println("Error Writing the len(reqPkt.Data)", err)
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, reqPkt.Data); err != nil {
		log.Println("Error Writing the reqPkt.Data", err)
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, reqPkt.Message); err != nil {
		log.Println("Error Writing the reqPkt.Message", err)
		return err
	}
	_, err := (*conn).Write(buf.Bytes())
	if err != nil {
		log.Println("Error Sending the request")
		return err
	}
	return nil
}

func ReadResponse(conn *net.Conn, respPkt *RespPacket) error {
	buf := make([]byte, 1024)
	n, err := (*conn).Read(buf)
	if err != nil {
		if err == io.EOF {
			log.Println("Connection closed")
			return err
		}
		log.Println("Error reading the request:", err)
		return err
	}
	dataBuf := bytes.NewReader(buf[:n])
	if err := binary.Read(dataBuf, binary.BigEndian, &respPkt.RespType); err != nil {
		log.Println("Error Reading the respPkt.Status", err)
		return err
	}
	if err := binary.Read(dataBuf, binary.BigEndian, &respPkt.Result); err != nil {
		log.Println("Error Reading the respPkt.Result", err)
		return err
	}
	if err := binary.Read(dataBuf, binary.BigEndian, &respPkt.Status); err != nil {
		log.Println("Error Reading the respPkt.Status", err)
		return err
	}
	messageBytes := make([]byte, dataBuf.Len())
	if err := binary.Read(dataBuf, binary.BigEndian, messageBytes); err != nil {
		log.Println("Error Reading the respPkt.Message", err)
		return err
	}
	respPkt.Message = []byte(messageBytes)
	return nil
}

func WriteResponse(conn *net.Conn, respPkt RespPacket) error {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, respPkt.RespType); err != nil {
		log.Println("Error Writing the reqPkt.ReqType", err)
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, respPkt.Result); err != nil {
		log.Println("Error Writing the reqPkt.Result", err)
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, respPkt.Status); err != nil {
		log.Println("Error Writing the reqPkt.Data", err)
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, respPkt.Message); err != nil {
		log.Println("Error Writing the reqPkt.Message", err)
		return err
	}
	_, err := (*conn).Write(buf.Bytes())
	if err != nil {
		log.Println("Error Sending the Response")
		return err
	}
	return nil
}
