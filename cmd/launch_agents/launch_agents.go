package main

import (
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent/server"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	myServer := server.NewServer("TEST", "127.0.0.1:8080")
	go myServer.Start()

	time.Sleep(time.Second)

	var hobbies []string = []string{"football", "game", "paint", "horror", "manga", "history"}
	var agts_m []*agt.AgentManager
	var agts_w []*agt.AgentWorker

	id_m := 0
	id_w := 0

	for i := 0; i < 100; i++ {
		if rand.Intn(2) == -1 {
			id := "agt_m" + strconv.Itoa(id_m)
			agts_m = append(agts_m, agt.NewAgentManager(id, hobbies[rand.Intn(6)], (*agt.Server)(myServer)))
			id_m++
		} else {
			id := "agt_w" + strconv.Itoa(id_w)
			agts_w = append(agts_w, agt.NewAgentWorker(id, agt.MakeRandomSliceOfHobbies(hobbies), (*agt.Server)(myServer)))
			id_w++
		}
	}

	id := "agt_m" + strconv.Itoa(id_m)
	agts_m = append(agts_m, agt.NewAgentManager(id, hobbies[rand.Intn(6)], (*agt.Server)(myServer)))
	id_m++

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
