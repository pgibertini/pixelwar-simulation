package agent

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewServer(id string, addr string) *Server {
	cin := make(chan (interface{}))
	return &Server{
		identifier: id,
		address:    addr,
		Cin:        cin,
		// TODO: add a slices a Canvas (like in the vote API)
	}
}

func (srv *Server) Start() {
	// Multiplexage des différentes requêtes possibles
	mux := http.NewServeMux()
	mux.HandleFunc("/test_request", srv.makeTestRequest)

	// TODO: add "newCanvas" function
	// TODO: add "paintPixel" function

	// Création d'un serveur web
	s := &http.Server{
		Addr:           srv.address,
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Launching REST server:", srv.address)
	go func() {
		for {
			value := <-srv.Cin
			switch value.(type) {
			case *AgentManager:
				srv.registerManager(value.(*AgentManager))
			case *AgentWorker:
				srv.registerWorker(value.(*AgentWorker))
			default:
				fmt.Println("Error: bad request")
			}
		}
	}()
	go log.Fatal(s.ListenAndServe())
}

func (srv *Server) registerManager(am *AgentManager) {
	fmt.Printf("Registering a manager. ID = %s\n", am.id)
	srv.ams = append(srv.ams, am)
}

func (srv *Server) registerWorker(aw *AgentWorker) {
	fmt.Printf("Registering a worker. ID = %s\n", aw.id)
	srv.aws = append(srv.aws, aw)
}

func (srv *Server) GetManagers() []*AgentManager {
	return srv.ams
}

func (srv *Server) GetWorkers() []*AgentWorker {
	return srv.aws
}
