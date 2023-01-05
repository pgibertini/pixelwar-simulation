package agent

import (
	"fmt"
	"log"
)

func NewChat() *Chat {
	cin := make(chan (interface{}))
	return &Chat{
		Cin: cin,
	}
}

func (srv *Chat) Start() {
	go func() {
		for {
			value := <-srv.Cin
			switch value.(type) {
			case *AgentManager:
				srv.registerManager(value.(*AgentManager))
			case *AgentWorker:
				srv.registerWorker(value.(*AgentWorker))
			case FindWorkersRequest:
				srv.findWorkersRespond(value.(FindWorkersRequest))
				(srv.Ams[GetManagerIndex(srv.Ams, (value.(FindWorkersRequest)).Id_manager)]).C_findWorkers <- srv.findWorkersRespond(value.(FindWorkersRequest))
			default:
				fmt.Println("Error: bad request")
			}
		}
	}()

	go log.Println("Serveur de chat lancé")
}

func (srv *Chat) registerManager(am *AgentManager) {
	fmt.Printf("Registering a manager. ID = %s\n", am.GetID())
	srv.Ams = append(srv.Ams, am)
}

func (srv *Chat) registerWorker(aw *AgentWorker) {
	fmt.Printf("Registering a worker. ID = %s\n", aw.GetID())
	srv.Aws = append(srv.Aws, aw)
}

func (srv *Chat) findWorkersRespond(req FindWorkersRequest) FindWorkersResponse {
	var resp FindWorkersResponse
	for _, v := range srv.Aws {
		if ContainsHobby(v.Hobbies, req.Hobby) {
			resp.Workers = append(resp.Workers, v)
		}
	}
	return resp
}

func (srv *Chat) GetManagers() []*AgentManager {
	return srv.Ams
}

func (srv *Chat) GetWorkers() []*AgentWorker {
	return srv.Aws
}
