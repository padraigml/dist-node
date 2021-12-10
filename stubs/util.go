package stubs

import "net/rpc"

type Event struct{}

type Info struct {
	StartY int
	EndY   int
	Turn   int
	World  [][]byte
	P      Params
}

type AliveInfo struct {
	World    [][]byte
	NumAlive int
}

type NodeInfo struct {
	Err   error
	World [][]byte
}

type NodeStruct struct {
	Node         *rpc.Client
	ShuttingDown bool
}

type Job struct {
	StartY    int
	EndY      int
	TurnCount int
	World     [][]byte
	P         Params
}

type SaveWorldInfo struct {
	World [][]byte
	Turn  int
}

type Params struct {
	Turns       int
	Threads     int
	ImageWidth  int
	ImageHeight int
}
