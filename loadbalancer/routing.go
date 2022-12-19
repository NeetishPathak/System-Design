package main

import "log"

const (
	algoRR   = "round-robin"
	algoWRR  = "weighted-round-robin"
	algoConn = "leastConn"
)

type IRoutingAlgorithm interface {
	getName() string
	getBackend([]*Backend) *Backend
	setNextBackend([]*Backend)
}

type RoundRobin struct {
	name  string
	index int
}

func (rr RoundRobin) getName() string {
	return rr.name
}

func (rr RoundRobin) getBackend(backends []*Backend) *Backend {
	if len(backends) < 1 {
		log.Panicln("No backends are configured")
	}
	return backends[rr.index]
}

func (rr *RoundRobin) setNextBackend(backends []*Backend) {
	if len(backends) < 1 {
		log.Panicln("No backends are configured")
	}
	rr.index = (rr.index + 1) % (len(backends))
}

type WeightedRoundRobin struct {
	name          string
	index         int
	floatingIndex float32
}

func (rr WeightedRoundRobin) getName() string {
	return rr.name
}

func (rr WeightedRoundRobin) getBackend(backends []*Backend) *Backend {
	if len(backends) < 1 {
		log.Panicln("No backends are configured")
	}
	return backends[rr.index]
}

func (rr *WeightedRoundRobin) setNextBackend(backends []*Backend) {
	if len(backends) < 1 {
		log.Panicln("No backends are configured")
	}
	rr.floatingIndex = rr.floatingIndex + (1.0 / float32((*backends[rr.index]).ratio))
	floor_floatingIndex := int(rr.floatingIndex)
	rr.index = (floor_floatingIndex) % (len(backends))
}
