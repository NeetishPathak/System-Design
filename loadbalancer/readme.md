# Toy Load balancer

## Description 
- Create a TCP Loadbalancer to route requests to the BackEnds
- Backend can be enabled disabled by an external user
- Backend selection Algorithms 
  - Round Robin
  - Weighted Round Robin
  - Connection Count Based
- Backend will be a TCP server , that echos the client message with the server name appended 

### Usage
```bash
Usage: go run . -a <lb routing algorithm>
lb_routing_algo
rr : Round Robin
wrr : Weighted Round Robin
lc : Least Connection
```

### Round Robin 

1. Start Backends and Round Robin LB
```bash
go run . -a rr
Backed Info - Name: B1, Endpoint: 127.0.0.1:25601, Enabled:true, Ratio: 2
Backed Info - Name: B2, Endpoint: 127.0.0.1:25602, Enabled:true, Ratio: 3
Backed Info - Name: B3, Endpoint: 127.0.0.1:25603, Enabled:true, Ratio: 4
LB Info - Ip: 127.0.0.1 , Port: 8080, Algo: round-robin
```

2. Run client test program
```bash
go test -v -run TestDial
=== RUN   TestDial
Client Connected Successfully {1 Android}
Snt:  Hello Android
Rcv:  B1-Hello Android

Client Connected Successfully {2 Iphone}
Snt:  Hello Iphone
Rcv:  B2-Hello Iphone

Client Connected Successfully {3 Windows}
Snt:  Hello Windows
Rcv:  B3-Hello Windows

Client Connected Successfully {4 MacOS}
Snt:  Hello MacOS
Rcv:  B1-Hello MacOS

Client Connected Successfully {5 Ubuntu}
Snt:  Hello Ubuntu
Rcv:  B2-Hello Ubuntu

Client Connected Successfully {6 SUSE}
Snt:  Hello SUSE
Rcv:  B3-Hello SUSE

Client Connected Successfully {7 Kali}
Snt:  Hello Kali
Rcv:  B1-Hello Kali

Client Connected Successfully {8 Fedora}
Snt:  Hello Fedora
Rcv:  B2-Hello Fedora

--- PASS: TestDial (0.01s)
PASS
ok      loadbalancer    0.248s
```

3. Logs on Server Side
```bash
go run . -a rr
Backed Info - Name: B1, Endpoint: 127.0.0.1:25601, Enabled:true, Ratio: 2
Backed Info - Name: B2, Endpoint: 127.0.0.1:25602, Enabled:true, Ratio: 3
Backed Info - Name: B3, Endpoint: 127.0.0.1:25603, Enabled:true, Ratio: 4
LB Info - Ip: 127.0.0.1 , Port: 8080, Algo: round-robin
2022/12/19 00:00:07 Loadbalancer connected to:  127.0.0.1:61563
2022/12/19 00:00:07 B1  connected to:  127.0.0.1:61564
2022/12/19 00:00:07 Loadbalancer connected to:  127.0.0.1:61565
2022/12/19 00:00:07 B2  connected to:  127.0.0.1:61566
2022/12/19 00:00:07 Loadbalancer connected to:  127.0.0.1:61567
2022/12/19 00:00:07 B3  connected to:  127.0.0.1:61568
2022/12/19 00:00:07 Loadbalancer connected to:  127.0.0.1:61569
2022/12/19 00:00:07 B1  connected to:  127.0.0.1:61570
2022/12/19 00:00:07 Loadbalancer connected to:  127.0.0.1:61571
2022/12/19 00:00:07 B2  connected to:  127.0.0.1:61572
2022/12/19 00:00:07 Loadbalancer connected to:  127.0.0.1:61573
2022/12/19 00:00:07 B3  connected to:  127.0.0.1:61574
2022/12/19 00:00:07 Loadbalancer connected to:  127.0.0.1:61575
2022/12/19 00:00:07 B1  connected to:  127.0.0.1:61576
2022/12/19 00:00:07 Loadbalancer connected to:  127.0.0.1:61577
2022/12/19 00:00:07 B2  connected to:  127.0.0.1:61578
```

### Weighted Round Robin 

1. Start Backends and Round Robin LB
```bash
go run . -a wrr
Backed Info - Name: B1, Endpoint: 127.0.0.1:25601, Enabled:true, Ratio: 2
Backed Info - Name: B2, Endpoint: 127.0.0.1:25602, Enabled:true, Ratio: 3
Backed Info - Name: B3, Endpoint: 127.0.0.1:25603, Enabled:true, Ratio: 4
LB Info - Ip: 127.0.0.1 , Port: 8080, Algo: weighted-round-robin
```

2. Run client test program

The requests are served by different backends based on the weights (ratio) assigned to them during intialization

```bash
go test -v -run TestDial
=== RUN   TestDial
Client Connected Successfully {1 Android}
Snt:  Hello Android
Rcv:  B1-Hello Android

Client Connected Successfully {2 Iphone}
Snt:  Hello Iphone
Rcv:  B1-Hello Iphone

Client Connected Successfully {3 Windows}
Snt:  Hello Windows
Rcv:  B2-Hello Windows

Client Connected Successfully {4 MacOS}
Snt:  Hello MacOS
Rcv:  B2-Hello MacOS

Client Connected Successfully {5 Ubuntu}
Snt:  Hello Ubuntu
Rcv:  B2-Hello Ubuntu

Client Connected Successfully {6 SUSE}
Snt:  Hello SUSE
Rcv:  B3-Hello SUSE

Client Connected Successfully {7 Kali}
Snt:  Hello Kali
Rcv:  B3-Hello Kali

Client Connected Successfully {8 Fedora}
Snt:  Hello Fedora
Rcv:  B3-Hello Fedora

--- PASS: TestDial (0.01s)
PASS
ok      loadbalancer    0.153s
```

3. Logs on Server Side
```bash
go run . -a wrr
Backed Info - Name: B1, Endpoint: 127.0.0.1:25601, Enabled:true, Ratio: 2
Backed Info - Name: B2, Endpoint: 127.0.0.1:25602, Enabled:true, Ratio: 3
Backed Info - Name: B3, Endpoint: 127.0.0.1:25603, Enabled:true, Ratio: 4
LB Info - Ip: 127.0.0.1 , Port: 8080, Algo: weighted-round-robin
2022/12/19 00:01:20 Loadbalancer connected to:  127.0.0.1:61624
2022/12/19 00:01:20 B1  connected to:  127.0.0.1:61625
2022/12/19 00:01:20 Loadbalancer connected to:  127.0.0.1:61626
2022/12/19 00:01:20 B1  connected to:  127.0.0.1:61627
2022/12/19 00:01:20 Loadbalancer connected to:  127.0.0.1:61628
2022/12/19 00:01:20 B2  connected to:  127.0.0.1:61629
2022/12/19 00:01:20 Loadbalancer connected to:  127.0.0.1:61630
2022/12/19 00:01:20 B2  connected to:  127.0.0.1:61631
2022/12/19 00:01:20 Loadbalancer connected to:  127.0.0.1:61632
2022/12/19 00:01:20 B2  connected to:  127.0.0.1:61633
2022/12/19 00:01:20 Loadbalancer connected to:  127.0.0.1:61634
2022/12/19 00:01:20 B3  connected to:  127.0.0.1:61635
2022/12/19 00:01:20 Loadbalancer connected to:  127.0.0.1:61636
2022/12/19 00:01:20 B3  connected to:  127.0.0.1:61637
2022/12/19 00:01:20 Loadbalancer connected to:  127.0.0.1:61638
2022/12/19 00:01:20 B3  connected to:  127.0.0.1:61639
```



