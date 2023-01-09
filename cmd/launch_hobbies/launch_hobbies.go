package Hobbies

import (
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func LaunchHobbies(random bool, propo bool, nbAgents int, cooldown int, size int) {
	// PARAMETERS
	url := "http://localhost:5555"
	var hobbies = []string{"2b2t", "Asterix", "Avengers", "BlueMario", "Canada", "ChronoTrigger", "CloneHero",
		"DeadCells", "FireEmblem", "France", "Hytale", "Kirby", "Linux", "LofiGirl", "Mario", "MTG", "NBA", "NecoArc",
		"OnePiece", "StarWars", "Technoblade"}

	minWorkers := 5
	maxWorkers := 10

	hobbyWorkerMap := map[string]struct {
		nWorkers int
		floatVal float64
	}{}

	for _, h := range hobbies {
		var nWorkers int
		if random { // generate a random number between 5 and 200
			nWorkers = rand.Intn(maxWorkers-minWorkers+1) + minWorkers
		} else if propo { //use a number proportional to the size of the image
			//nWorkers = getNbWorkers(h)
			nWorkers = nbAgents
		} else { //use a fixed number
			nWorkers = nbAgents
		}
		floatVal := 3 + rand.Float64()*5
		hobbyWorkerMap[h] = struct {
			nWorkers int
			floatVal float64
		}{nWorkers, floatVal}
	}

	rand.Seed(time.Now().UnixNano())

	// create a new place
	placeID := agt.CreateNewPlace(url, size, size, cooldown)
	log.Printf("Playing on %s", placeID)

	// chat for agents discussion
	myChat := agt.NewChat(placeID, url, cooldown, size, size)
	go myChat.Start()

	var managers []*agt.AgentManager
	var workers []*agt.AgentWorker

	// initializing managers and workers
	for _, h := range hobbies {
		nWorkers := hobbyWorkerMap[h].nWorkers
		conquestVal := hobbyWorkerMap[h].floatVal
		managers = append(managers, agt.NewAgentManager(h, h, conquestVal, myChat))

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

/*func getNbWorkers(hobby string) (nWorkers int) {

}*/
