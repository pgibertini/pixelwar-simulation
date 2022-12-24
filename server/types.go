package server

import (
	"sync"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
)

type Server struct {
	sync.Mutex
	identifier string
	address    string
	ams        []*agent.AgentManager
	aws        []*agent.AgentWorker
}

type testRequest struct {
	Value string `json:"value"`
}
