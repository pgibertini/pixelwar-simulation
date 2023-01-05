package agent

import (
	"gitlab.utc.fr/pixelwar_ia04/pixelwar/painting"
	"sync"
	"time"
)

type Server struct {
	sync.Mutex
	Identifier string
	Address    string
	Places     map[string]*Place
	Ams        []*AgentManager
	Aws        []*AgentWorker
	Cin        chan interface{}
}

type Place struct {
	Id         string
	Canvas     *painting.CanvasHex
	LastAction map[string]time.Time
	Cooldown   time.Duration
}

type Chat struct {
	Ams []*AgentManager
	Aws []*AgentWorker
	Cin chan interface{}
}

type AgentWorker struct {
	id      string
	tab     []painting.PixelToPlace
	Hobbies []string
	Cout    chan interface{}
	Srv     *Server
}

type AgentManager struct {
	id              string
	agts            []*AgentWorker
	hobby           string
	Srv             *Server
	bufferImgLayout []painting.PixelToPlace
	Cin             chan interface{}
	C_findWorkers   chan FindWorkersResponse
}
