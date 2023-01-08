package main

import (
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"log"
	"strconv"
)

func main() {
	// PARAMETERS
	url := "http://localhost:8080"
	nWorkers := 10
	size := 1000

	// create a new place
	placeID := agt.CreateNewPlace(url, size, size)
	log.Printf("Playing on %s", placeID)

	// chat for agents discussion
	myChat := agt.NewChat()
	go myChat.Start()

	var hobbies = []string{"#FF0000", "#00FF00", "#0000FF"}
	var managers []*agt.AgentManager
	var workers []*agt.AgentWorker

	// initializing managers and workers
	for i, h := range hobbies {
		managers = append(managers, agt.NewAgentManager(strconv.Itoa(i), h, myChat, placeID, url))

		for j := 0; j < nWorkers; j++ {
			workers = append(workers, agt.NewAgentWorker(h+strconv.Itoa(j), []string{h}, myChat, placeID, url))
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
	var pixels []painting.HexPixel
	for _, m := range managers {
		pixels = nil
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				pixels = append(pixels, painting.HexPixel{X: i, Y: j, Color: painting.HexColor(m.GetHobby())})
			}
		}
		painting.ShuffleHexPixels(pixels)
		m.AddPixelsToPlace(pixels)
	}

	// sending pixel to workers
	for _, m := range managers {
		m.DistributeWork()
	}

	fmt.Scanln()
}
