package agent

import "sync"

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

type findWorkersRequest struct {
	Id_manager string
	hobby      string
}

type findWorkersResponse struct {
	workers []*AgentWorker
}
