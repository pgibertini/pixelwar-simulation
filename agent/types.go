package agent

import (
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"sync"
)

type Chat struct {
	Ams      []*AgentManager
	Aws      []*AgentWorker
	Cin      chan interface{}
	srvUrl   string
	placeId  string
	cooldown int
}

type AgentWorker struct {
	id      string
	tab     []painting.HexPixel
	Hobbies []string
	Cin     chan interface{}
	Chat    *Chat
	mu      sync.Mutex
}

type AgentManager struct {
	id            string
	workers       []*AgentWorker
	hobby         string
	Chat          *Chat
	Painting      painting.ManagerPainting
	ImgLayout     [][]painting.HexColor
	pixelsToPlace []painting.HexPixel
	Cout          chan interface{}
	Cin           chan FindWorkersResponse
}
