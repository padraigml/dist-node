package stubs

var RegisterNode = "Broker.RegisterNode"

type BrokerRequest struct {
	Info Info
	Key  int
}

type NResponse struct {
	Inf      NodeInfo
	NumAlive int
}

type SaveResponse struct {
	Info SaveWorldInfo
}

type PauseResponse struct {
	Turn   int
	Paused bool
}

type PublishRequest struct {
	Job Job
}

type NodeResponse struct {
	Info     Info
	Turn     int
	NumAlive int
	Ready    bool
}

type Subscription struct {
	NodeAddress string
	Callback    string
}
