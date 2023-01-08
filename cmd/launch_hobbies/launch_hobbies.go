package main

import (
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// PARAMETERS
	url := "http://localhost:8080"
	nWorkers := 100
	size := 500
	cooldown := 1

	// create a new place
	placeID := agt.CreateNewPlace(url, size, size, cooldown)
	log.Printf("Playing on %s", placeID)

	// chat for agents discussion
	myChat := agt.NewChat()
	go myChat.Start()

	var hobbies = []string{"2b2t", "Asterix", "Avengers", "BlueMario", "Canada", "ChronoTrigger"}
	var managers []*agt.AgentManager
	var workers []*agt.AgentWorker

	// initializing managers and workers
	for i, h := range hobbies {
		managers = append(managers, agt.NewAgentManager(strconv.Itoa(i), h, myChat, placeID, url))

		for j := 0; j < nWorkers; j++ {
			workers = append(workers, agt.NewAgentWorker(h+strconv.Itoa(j), []string{h}, myChat, placeID, url, cooldown))
		}
	}

	// starting the agents
	for _, w := range workers {
		w.Start()
	}

	for _, m := range managers {
		m.Start()
	}

	// giving pixel to place
	for _, m := range managers {
		rand.Seed(time.Now().UnixNano())
		m.LoadLayoutFromFile(fmt.Sprintf("./images/%s", m.GetHobby()))
		m.AddPixelsToPlace(painting.ImgLayoutToPixelList(m.ImgLayout, rand.Intn(size/2), rand.Intn(size/2)))
	}

	// sending pixel to workers
	for _, m := range managers {
		go m.DistributeWork()
	}

	fmt.Scanln()
}
