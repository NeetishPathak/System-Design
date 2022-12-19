package main

import (
	"flag"
	"fmt"
)

func cli_help() {
	msg := fmt.Sprintln(`Usage: go run . -a <lb routing algorithm>
lb_routing_algo
rr : Round Robin
wrr : Weighted Round Robin
lc : Least Connection`)
	fmt.Print(msg)
}
func main() {
	var lbAlgo string
	flag.StringVar(&lbAlgo, "a", "rr", "LB Algo Type")
	flag.Parse()

	if lbAlgo == "rr" || lbAlgo == "wrr" || lbAlgo == "lc" {
		var myLb ILoadBalancer
		myBackEnds := []*Backend{
			{"127.0.0.1", "25601", "B1", true, 2},
			{"127.0.0.1", "25602", "B2", true, 3},
			{"127.0.0.1", "25603", "B3", true, 4},
		}
		for _, bk := range myBackEnds {
			go bk.Initialize()
			bk.printBackend()
		}
		myLb = &TcpLoadBalancer{"127.0.0.1", "8080", myBackEnds, nil}

		switch lbAlgo {
		case "rr":
			myLb.initialize(algoRR)
		case "wrr":
			myLb.initialize(algoWRR)
		case "lc":
			myLb.initialize(algoConn)
		default:
			myLb.initialize(algoRR)
		}
	} else {
		cli_help()
	}
}
