package main

import (
	"fmt"
	"time"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/server"
)

func main() {

	myServer := server.NewServer("TEST", "127.0.0.1:8080")
	go myServer.Start()

	time.Sleep(5 * time.Second)

	var agts []*agent.AgentManager

	myAgent := agent.NewAgentManager("ag1_m", "football")

	agts = append(agts, myAgent)

	go myAgent.Start()

	fmt.Scanln()

	fmt.Println(myAgent)
	fmt.Println(myServer.GetManagers()[0])
}
