package main

import (
	"fmt"
	"time"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
)

func main() {

	myServer := agent.NewServer("TEST", "127.0.0.1:8080")
	go myServer.Start()

	time.Sleep(5 * time.Second)

	var agts []*agent.AgentManager

	myAgent := agent.NewAgentManager("ag1_m", "football", myServer)
	myAgent1 := agent.NewAgentManager("ag2_m", "football", myServer)
	myAgent8 := agent.NewAgentManager("ag2_m", "football", myServer)
	myAgent5 := agent.NewAgentWorker("ag5_w", nil, myServer)
	myAgent4 := agent.NewAgentManager("ag4_m", "football", myServer)
	myAgent2 := agent.NewAgentWorker("ag2_w", nil, myServer)
	myAgent3 := agent.NewAgentManager("ag4_m", "football", myServer)

	agts = append(agts, myAgent)

	go myAgent.Start()

	go myAgent1.Start()

	go myAgent2.Start()

	go myAgent3.Start()

	go myAgent5.Start()

	go myAgent4.Start()

	go myAgent8.Start()

	fmt.Scanln()

	man := myServer.GetManagers()

	for _, value := range man {
		fmt.Println(*value)
	}
}
