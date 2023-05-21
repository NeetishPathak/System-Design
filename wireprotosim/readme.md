# Wire Protocol

### A Toy wire protocol

This basic wire protocol showcases an example interaction between a client and server.

- We define a ReqPacket and RespPacket 

ReqPacket

```golang
type ReqPacket struct {
	ReqType int32
	Data    []int32
	Message []byte
}
```

When sending the request, the buffer is serialized as byte slice of variable length depending on length of Data and Message

```

      |0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|
      |              0|              1|              2|              3|
------+---------------+---------------+---------------+---------------+
    0 | ReqType                                                       |
------+---------------+---------------+-------------------------------+
    4 | Length of Data   (Example : 2 data values are passed)         |                                             
------+---------------+---------------+-------------------------------+
    8 | Data[0]                                                       |
------+---------------+---------------+-------------------------------+
   12 | Data[1]                                                       |
------+---------------+---------------+-------------------------------+
   16 | Message[:4]                                                   |    
------+---------------------------------------------------------------+
   20 | Message[:8]                                                   |    
------+---------------------------------------------------------------+
   24 | Message...                                                    |    
------+---------------------------------------------------------------+

```

Response Packet

```golang
type RespPacket struct {
	RespType int32
	Result   int32
	Status   int32
	Message  []byte
}
```

The response buffer is serialized as binary encoded byte slice of variable length depending on Message size

```
      |0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|0|1|2|3|4|5|6|7|
      |              0|              1|              2|              3|
------+---------------+---------------+---------------+---------------+
    0 | RespType                                                      |
------+---------------+---------------+-------------------------------+
    4 | Result                                                        |
------+---------------+---------------+-------------------------------+
    8 | Status                                                        |
------+---------------+---------------+-------------------------------+
    12 | Message...                                                   |                                             
------+---------------+---------------+-------------------------------+
```

- Handles two types of requests

```golang
const (
	PingReq = 1
	AddReq  = 2
)

```

### Watch a demo
https://github.com/NeetishPathak/System-Design/assets/4778360/21b9ccec-c277-466f-b3d9-8a4c9faa9dd4


