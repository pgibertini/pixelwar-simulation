package agent

import (
	"sync"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
)

type Server struct {
	sync.Mutex
	identifier string
	address    string
	ams        []*AgentManager
	aws        []*AgentWorker
	Cin        chan interface{}
}

type testRequest struct {
	Value string `json:"value"`
}

type sendPixelsRequest struct {
	pixels []painting.PixelToPlace
	id_am  string
}

type findWorkersRequest struct {
	Id_manager string
	hobby      string
}

type findWorkersResponse struct {
	workers []*AgentWorker
}
