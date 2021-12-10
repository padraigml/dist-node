package main

var RegisterNode = "Broker.RegisterNode"

type BrokerRequest struct {
	Info structs.Info
	Key  int
}

type NResponse struct {
	Inf      structs.NodeInfo
	NumAlive int
}

type SaveResponse struct {
	Info structs.SaveWorldInfo
}

type PauseResponse struct {
	Turn   int
	Paused bool
}

type PublishRequest struct {
	Job structs.Job
}

type NodeResponse struct {
	Info     structs.Info
	Turn     int
	NumAlive int
	Ready    bool
}

type Subscription struct {
	NodeAddress string
	Callback    string
}
