package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
)

type AgentWorker struct {
	id      string
	tab     []painting.PixelToPlace
	hobbies []string
	Cout    chan interface{}
}

type AgentManager struct {
	id    string
	agts  []string
	hobby string
}

func NewAgentWorker(idAgt string, hobbiesAgt []string) *AgentWorker {
	channel := make(chan interface{})
	return &AgentWorker{
		id:      idAgt,
		hobbies: hobbiesAgt,
		Cout:    channel,
	}
}

func (aw *AgentWorker) Start() {

}

func (aw *AgentWorker) GetID() string {
	return aw.id
}

func (aw *AgentWorker) GetAddress() **AgentWorker {
	return &aw
}

func (aw *AgentWorker) getManagers(hobby string) {

}

func (aw *AgentWorker) drawOnePixel(pixel painting.Pixel) {

}

// ============================ AgentManager ============================

func NewAgentManager(idAgt string, hobbyAgt string) *AgentManager {
	return &AgentManager{
		id:    idAgt,
		hobby: hobbyAgt,
	}
}

func (am *AgentManager) Start() {
	err := am.register()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("OK")
}

func (am *AgentManager) GetID() string {
	return am.id
}

func (am *AgentManager) GetAddress() *AgentManager {
	return am
}

func (am *AgentManager) register() (err error) {
	req := RegisterAMRequest{am}

	url := "http://127.0.0.1:8080/register_am"
	data, _ := json.Marshal(req)

	resp, err := http.Post(url, "applications/json", bytes.NewBuffer(data))

	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Error")
	}

	return
}
