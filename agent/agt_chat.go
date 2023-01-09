package agent

import (
	"flag"
	"fmt"
	"log"
)

var Debug bool

func init() {
	flag.BoolVar(&Debug, "debug-agt", false, "enable debug mode")
}

func NewChat(placeID string, url string, cooldown, height, width int) *Chat {
	cin := make(chan interface{})
	return &Chat{
		Cin:      cin,
		srvUrl:   url,
		placeId:  placeID,
		cooldown: cooldown,
		height:   height,
		width:    width,
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
				go func(req FindWorkersRequest) {
					resp := srv.findWorkersRespond(req)
					(srv.Ams[GetManagerIndex(srv.Ams, req.IdManager)]).Cin <- resp
				}(value.(FindWorkersRequest))
			default:
				fmt.Println("Error: bad request")
			}
		}
	}()

	go log.Println("Serveur de chat lancé")
}

func (srv *Chat) registerManager(am *AgentManager) {
	if Debug {
		log.Printf("Registering a manager: ID=%s\n", am.GetID())
	}
	srv.Ams = append(srv.Ams, am)
}

func (srv *Chat) registerWorker(aw *AgentWorker) {
	if Debug {
		log.Printf("Registering a worker: ID=%s\n", aw.GetID())
	}
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

// GETTERS

func (srv *Chat) GetManagers() []*AgentManager {
	return srv.Ams
}

func (srv *Chat) GetWorkers() []*AgentWorker {
	return srv.Aws
}

func (srv *Chat) GetURL() string {
	return srv.srvUrl
}

func (srv *Chat) GetPlaceID() string {
	return srv.placeId
}

func (srv *Chat) GetCooldown() int {
	return srv.cooldown
}

func (srv *Chat) GetHeight() int {
	return srv.height
}

func (srv *Chat) GetWidth() int {
	return srv.width
}
