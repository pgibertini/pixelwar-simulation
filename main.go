package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	myServer := agent.NewServer("TEST", "127.0.0.1:8080")
	go myServer.Start()

	time.Sleep(5 * time.Second)

	var hobbies []string = []string{"football", "game", "paint", "horror", "manga", "history"}
	var agts_m []*agent.AgentManager
	var agts_w []*agent.AgentWorker

	id_m := 0
	id_w := 0

	for i := 0; i < 10; i++ {
		if rand.Intn(2) == 0 {
			id := "agt_m" + strconv.Itoa(id_m)
			agts_m = append(agts_m, agent.NewAgentManager(id, hobbies[rand.Intn(6)], myServer))
			id_m++
		} else {
			id := "agt_w" + strconv.Itoa(id_w)
			agts_w = append(agts_w, agent.NewAgentWorker(id, hobbies, myServer))
			id_w++
		}
	}

	for _, v := range agts_w {
		v.Start()
	}

	for _, v := range agts_m {
		v.Start()
	}

	fmt.Scanln()

	man := myServer.GetManagers()
	wor := myServer.GetWorkers()

	for _, value := range man {
		fmt.Println(*value)
	}
	for _, value := range wor {
		fmt.Println(*value)
	}
}
