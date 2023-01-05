package agent

import (
	"fmt"
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
)

func NewAgentWorker(idAgt string, hobbiesAgt []string, srv *Server) *AgentWorker {
	channel := make(chan interface{})
	return &AgentWorker{
		id:      idAgt,
		Hobbies: hobbiesAgt,
		Cout:    channel,
		Srv:     srv}
}

func (aw *AgentWorker) Start() {
	aw.register()

	go func() {
		for {
			value := <-aw.Cout
			switch value.(type) {
			case sendPixelsRequest:
				aw.tab = append(aw.tab, value.(sendPixelsRequest).pixels...)
			default:
				fmt.Println("Error: bad request")
			}
		}
	}()
}

func (aw *AgentWorker) GetID() string {
	return aw.id
}

func (aw *AgentWorker) GetHobbies() []string {
	return aw.Hobbies
}

func (aw *AgentWorker) drawOnePixel(pixel painting.Pixel) {

}

func (aw *AgentWorker) register() {
	(aw.Srv).Cin <- aw
}
