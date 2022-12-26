package agent

import (
	"fmt"

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
	id    string
	agts  []string
	hobby string
	Srv   *Server
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
	err := aw.register()

	if err != nil {
		fmt.Println(err)
	}
}

func (aw *AgentWorker) GetID() string {
	return aw.id
}

func (aw *AgentWorker) GetHobbies() []string {
	return aw.hobbies
}

func (aw *AgentWorker) drawOnePixel(pixel painting.Pixel) {

}

func (aw *AgentWorker) register() (err error) {
	srv := aw.Srv
	srv.Cin <- aw

	if err != nil {
		return
	}

	return
}

// ============================ AgentManager ============================

func NewAgentManager(idAgt string, hobbyAgt string, srv *Server) *AgentManager {
	return &AgentManager{
		id:    idAgt,
		hobby: hobbyAgt,
		Srv:   srv}
}

func (am *AgentManager) Start() {
	err := am.register()

	if err != nil {
		fmt.Println(err)
	}
}

func (am *AgentManager) GetID() string {
	return am.id
}

func (am *AgentManager) register() (err error) {
	srv := am.Srv

	srv.Cin <- am

	if err != nil {
		return
	}

	return
}
