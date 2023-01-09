package launcher

import (
	"bufio"
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func LaunchPixelWar(nbWorkerPerManager string, cooldown, size int, conquestValue float64, modeScript bool) (placeID string) {
	// PARAMETERS
	url := "http://localhost:5555"
	var hobbies = []string{"2b2t", "Asterix", "Avengers", "BlueMario", "Canada", "ChronoTrigger", "CloneHero",
		"DeadCells", "FireEmblem", "France", "Hytale", "Kirby", "Linux", "LofiGirl", "Mario", "MTG", "NBA", "NecoArc",
		"OnePiece", "StarWars", "Technoblade"}

	minWorkers := 5
	maxWorkers := 30

	hobbyWorkerMap := map[string]struct {
		nWorkers    int
		conquestVal float64
	}{}

	for _, h := range hobbies {
		var nWorkers int
		switch nbWorkerPerManager {
		case "random":
			nWorkers = rand.Intn(maxWorkers-minWorkers+1) + minWorkers
		case "proportional":
			nWorkers = getNbWorkers(h)
		default:
			n, err := strconv.Atoi(nbWorkerPerManager)
			if err != nil {
				log.Println("Incorrect nbWorkerPerManager", err)
				return
			} else {
				nWorkers = n
			}
		}

		hobbyWorkerMap[h] = struct {
			nWorkers    int
			conquestVal float64
		}{nWorkers, conquestValue}
	}

	rand.Seed(time.Now().UnixNano())

	// create a new place
	placeID = agt.CreateNewPlace(url, size, size, cooldown)
	log.Printf("Playing on %s", placeID)

	// chat for agents discussion
	myChat := agt.NewChat(placeID, url, cooldown, size, size)
	go myChat.Start()

	var managers []*agt.AgentManager
	var workers []*agt.AgentWorker

	// initializing managers and workers
	for _, h := range hobbies {
		nWorkers := hobbyWorkerMap[h].nWorkers
		conquestVal := hobbyWorkerMap[h].conquestVal
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

	if modeScript {
		fmt.Scanln()
	}
	return
}

func getNbWorkers(hobby string) (nWorkers int) {
	filePath := filepath.Join("images", hobby)
	f, err := os.Open(filePath)
	if err != nil {
		return 1
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	// Get the dimensions of the layout
	scanner.Scan()
	str := scanner.Text()
	height, err := strconv.Atoi(str)
	if err != nil {
		return 1
	}
	scanner.Scan()
	str = scanner.Text()
	width, err := strconv.Atoi(str)
	if err != nil {
		return 1
	}
	return max(1, width*height/50)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
