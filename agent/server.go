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
		places:     make(map[string]*Place), // TODO: maybe refactor to only have 1 place. This will affect paint_pixel and get_pixel request as the id of the place will no longer be necessary
		Cin:        cin,
	}
}

// checkMethod test a method
func (srv *Server) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

func (srv *Server) Start() {
	// Multiplexage des différentes requêtes possibles
	mux := http.NewServeMux()
	mux.HandleFunc("/new_place", srv.doNewPlace)
	mux.HandleFunc("/paint_pixel", srv.doPaintPixel)
	mux.HandleFunc("/get_pixel", srv.doGetPixel)
	// TODO: add a get canva request that return the whole grid

	// Création d'un serveur web
	s := &http.Server{
		Addr:           srv.address,
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Listening on", srv.address)

	go func() {
		for {
			value := <-srv.Cin
			switch value.(type) {
			case *AgentManager:
				srv.registerManager(value.(*AgentManager))
			case *AgentWorker:
				srv.registerWorker(value.(*AgentWorker))
			case findWorkersRequest:
				srv.findWorkersRespond(value.(findWorkersRequest))
				(srv.ams[getManagerIndex(srv.ams, (value.(findWorkersRequest)).Id_manager)]).C_findWorkers <- srv.findWorkersRespond(value.(findWorkersRequest))
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

func (srv *Server) findWorkersRespond(req findWorkersRequest) findWorkersResponse {
	var resp findWorkersResponse
	for _, v := range srv.aws {
		if containsHobby(v.hobbies, req.hobby) {
			resp.workers = append(resp.workers, v)
		}
	}
	return resp
}

func (srv *Server) GetManagers() []*AgentManager {
	return srv.ams
}

func (srv *Server) GetWorkers() []*AgentWorker {
	return srv.aws
}