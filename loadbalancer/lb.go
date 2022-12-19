package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type ILoadBalancer interface {
	initialize(string)
	acceptConn()
	proxyConn(net.Conn)
}

type TcpLoadBalancer struct {
	Ip       string
	Port     string
	Backends []*Backend
	algo     IRoutingAlgorithm
}

func (tlb *TcpLoadBalancer) initialize(rAlgo string) {
	switch rAlgo {
	case algoRR:
		tlb.algo = &RoundRobin{algoRR, 0}
	case algoWRR:
		tlb.algo = &WeightedRoundRobin{algoWRR, 0, 0.0}
	case algoConn:
		tlb.algo = &RoundRobin{algoRR, 0}
	default:
		tlb.algo = &RoundRobin{algoRR, 0}
	}
	fmt.Printf("LB Info - Ip: %s , Port: %s, Algo: %s\n", tlb.Ip, tlb.Port, tlb.algo.getName())
	tlb.acceptConn()
}

func (tlb *TcpLoadBalancer) acceptConn() {

	laddr, err := net.ResolveTCPAddr("tcp", tlb.Ip+":"+tlb.Port)
	if err != nil {
		log.Panicln(err)
	}
	listener, err := net.Listen("tcp", laddr.String())
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("failed to accept conn:", err)
			conn.Close()
			continue
		}
		log.Println("Loadbalancer connected to: ", conn.RemoteAddr())
		go tlb.proxyConn(conn)
	}
}

func (tlb *TcpLoadBalancer) proxyConn(conn net.Conn) {

	// get the next backend server
	backend := tlb.algo.getBackend(tlb.Backends)
	tlb.algo.setNextBackend(tlb.Backends)

	// create a connection to the backend server
	backendAddr, err := net.ResolveTCPAddr("tcp", backend.ip+":"+backend.port)
	if err != nil {
		log.Panicln(err)
	}

	bkconn, err := net.Dial("tcp", backendAddr.String())
	if err != nil {
		log.Panicln("bkconn is not available ", err)
		_, _ = conn.Write([]byte("backend connection failure"))
		conn.Close()
	}

	go io.Copy(bkconn, conn)
	go io.Copy(conn, bkconn)

}
