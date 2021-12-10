package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"node/stubs"
	"time"
)

type Node struct {
	Close bool
}

var (
	//port     string
	KillNode = make(chan struct{})
)

func CalculateNextState(job stubs.Job) [][]byte {
	imgHeight := job.P.ImageHeight
	imgWidth := job.P.ImageWidth
	world := job.World
	// Create an empty world to store the result in.
	nw := make([][]byte, job.P.ImageHeight)
	for x := 0; x < job.P.ImageHeight; x++ {
		nw[x] = make([]byte, job.P.ImageWidth)
	}

	// Loop through the world between the given bounds
	for y := job.StartY; y < job.EndY; y++ {
		for x := 0; x < imgWidth; x++ {
			// Calculate the number of alive neighbours for each cell.
			alive := (world[(y+imgHeight-1)%imgHeight][(x+imgWidth-1)%imgWidth] / 255) +
				(world[(y+imgHeight-1)%imgHeight][(x+imgWidth)%imgWidth] / 255) +
				(world[(y+imgHeight-1)%imgHeight][(x+imgWidth+1)%imgWidth] / 255) +
				(world[(y+imgHeight)%imgHeight][(x+imgWidth-1)%imgWidth] / 255) +
				(world[(y+imgHeight)%imgHeight][(x+imgWidth+1)%imgWidth] / 255) +
				(world[(y+imgHeight+1)%imgHeight][(x+imgWidth-1)%imgWidth] / 255) +
				(world[(y+imgHeight+1)%imgHeight][(x+imgWidth)%imgWidth] / 255) +
				(world[(y+imgHeight+1)%imgHeight][(x+imgWidth+1)%imgWidth] / 255)
			if world[y][x] == 255 {
				if alive < 2 || alive > 3 {
					nw[y][x] = 0
				}
				if alive == 2 || alive == 3 {
					nw[y][x] = 255
				}
			}
			if world[y][x] == 0 && alive == 3 {
				nw[y][x] = 255
			}
		}
	}
	return nw[job.StartY:job.EndY]
}

func (n *Node) ProcessTurn(req stubs.PublishRequest, res *stubs.NResponse) (err error) {
	alive := calculateNumAlive(req.Job.World, req.Job.P)
	res.Inf.World = CalculateNextState(req.Job)
	res.NumAlive = alive
	return
}

func (n *Node) StopNode(req stubs.BrokerRequest, res *stubs.NodeResponse) (err error) {
	n.Close = true
	close(KillNode)
	return
}

func calculateNumAlive(world [][]byte, p stubs.Params) int {
	count := 0

	for y := 0; y < p.ImageHeight; y++ {
		for x := 0; x < p.ImageWidth; x++ {
			if world[y][x] == 255 {
				count += 1
			}
		}
	}

	return count
}

func Listen(pAddr *string, node *Node) {
	rpc.Register(node)
	listener, _ := net.Listen("tcp", *pAddr)
	defer listener.Close()
	rpc.Accept(listener)
}

func main() {
	pAddr := flag.String("port", ":8001", "Port to listen on.")
	ipAddr := flag.String("ip", "127.0.0.1", "IP address.")
	flag.Parse()
	//port = *pAddr

	rand.Seed(time.Now().UnixNano())

	go Listen(pAddr, &Node{Close: false})

	brokerAddr := "0.0.0.0:8001"
	client, err := rpc.Dial("tcp", brokerAddr)
	fmt.Println("here")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	fmt.Println("connected to broker: ", client)
	defer client.Close()

	request := stubs.Subscription{NodeAddress: *ipAddr + *pAddr, Callback: "Node.ProcessTurn"}
	response := new(stubs.NodeResponse)
	err2 := client.Call(stubs.RegisterNode, request, response)
	if err2 != nil {
		log.Fatal(err2)
	}

	<-KillNode
}
