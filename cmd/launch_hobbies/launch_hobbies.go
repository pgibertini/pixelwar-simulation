package main

import (
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"log"
	"strconv"
	"time"
)

func main() {
	// PARAMETERS
	url := "http://localhost:5555"
	nWorkers := 100
	size := 500
	cooldown := 10

	// create a new place
	placeID := agt.CreateNewPlace(url, size, size, cooldown)
	log.Printf("Playing on %s", placeID)

	// chat for agents discussion
	myChat := agt.NewChat(placeID, url, cooldown, size, size)
	go myChat.Start()

	var hobbies = []string{"2b2t", "Asterix", "Avengers", "BlueMario", "Canada", "ChronoTrigger"}
	var managers []*agt.AgentManager
	var workers []*agt.AgentWorker

	// initializing managers and workers
	for i, h := range hobbies {
		managers = append(managers, agt.NewAgentManager(strconv.Itoa(i), h, myChat))

		for j := 0; j < nWorkers; j++ {
			workers = append(workers, agt.NewAgentWorker(h+strconv.Itoa(j), []string{h}, myChat))
		}
	}

	// starting the agents
	for _, w := range workers {
		go w.Start()
	}

	time.Sleep(time.Second)

	for _, m := range managers {
		go m.Start()
	}

	fmt.Scanln()
}
