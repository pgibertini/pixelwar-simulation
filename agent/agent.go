package agent

import (
	"fmt"
	"sync"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
)

type AgentWorker struct {
	id      string
	tab     []painting.PixelToPlace
	hobbies []string
	Cout    chan interface{}
	Srv     *Server
}

type AgentManager struct {
	id            string
	agts          []*AgentWorker
	hobby         string
	Srv           *Server
	Cin           chan interface{}
	C_findWorkers chan findWorkersResponse
}

func NewAgentWorker(idAgt string, hobbiesAgt []string, srv *Server) *AgentWorker {
	channel := make(chan interface{})
	return &AgentWorker{
		id:      idAgt,
		hobbies: hobbiesAgt,
		Cout:    channel,
		Srv:     srv}
}

func (aw *AgentWorker) Start() {
	aw.register()
}

func (aw *AgentWorker) GetID() string {
	return aw.id
}

func (aw *AgentWorker) GetHobbies() []string {
	return aw.hobbies
}

func (aw *AgentWorker) drawOnePixel(pixel painting.Pixel) {

}

func (aw *AgentWorker) register() {
	(aw.Srv).Cin <- aw
}

// ============================ AgentManager ============================

func NewAgentManager(idAgt string, hobbyAgt string, srv *Server) *AgentManager {
	cin := make(chan interface{})
	cout := make(chan findWorkersResponse)
	return &AgentManager{
		id:            idAgt,
		hobby:         hobbyAgt,
		Srv:           srv,
		Cin:           cin,
		C_findWorkers: cout}
}

func (am *AgentManager) Start() {
	am.register()
	am.updateWorkers()
}

func (am *AgentManager) GetID() string {
	return am.id
}

func (am *AgentManager) register() {
	(am.Srv).Cin <- am
}

func (am *AgentManager) updateWorkers() {
	// Not sure of this. Can AgentWorkers change their hobbies?
	for k, v := range am.agts {
		if !containsHobby(v.hobbies, am.hobby) {
			am.agts = remove(am.agts, k)
		}
	}
	req := findWorkersRequest{am.id, am.hobby}
	(am.Srv).Cin <- req

	fmt.Println("Voici ma liste de workers : ", am.agts)

	var wg sync.WaitGroup
	var resp findWorkersResponse

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("J'attends une réponse du serveur...")
		resp = <-am.C_findWorkers
	}()

	wg.Wait()
	fmt.Println("Réponse reçue ! : ", resp)

	for _, v := range resp.workers {
		if exists(am.agts, v.id) == -1 {
			am.agts = append(am.agts, v)
		}
	}

	fmt.Println("Voici ma liste finale de workers : ", am.agts)

}

// ============================ Utilities ============================

func containsHobby(hobbies []string, hobby string) bool {
	for _, v := range hobbies {
		if v == hobby {
			return true
		}
	}
	return false
}

func exists(s []*AgentWorker, id string) int {
	for k, v := range s {
		if v.id == id {
			return k
		}
	}
	return -1
}

func getManagerIndex(s []*AgentManager, id string) int {
	for k, v := range s {
		if v.id == id {
			return k
		}
	}
	return -1
}

func remove(s []*AgentWorker, i int) []*AgentWorker {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
