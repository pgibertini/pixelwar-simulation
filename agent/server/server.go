package server

import (
	"flag"
	"fmt"
	agt "gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
	"log"
	"net/http"
	"time"
)

var debug bool

type Server agt.Server

func init() {
	flag.BoolVar(&debug, "debug", true, "enable debug mode")
	// TODO: fix to have value passed by a flag
}

func NewServer(id string, addr string) *Server {
	cin := make(chan (interface{}))
	return &Server{
		Identifier: id,
		Address:    addr,
		Places:     make(map[string]*agt.Place),
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
	mux.HandleFunc("/get_canvas", srv.doGetCanvas)
	mux.HandleFunc("/canvas", srv.doCanvas)

	// Création d'un serveur web
	s := &http.Server{
		Addr:           srv.Address,
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Listening on", srv.Address)

	go func() {
		for {
			value := <-srv.Cin
			switch value.(type) {
			case *agt.AgentManager:
				srv.registerManager(value.(*agt.AgentManager))
			case *agt.AgentWorker:
				srv.registerWorker(value.(*agt.AgentWorker))
			case agt.FindWorkersRequest:
				srv.findWorkersRespond(value.(agt.FindWorkersRequest))
				(srv.Ams[agt.GetManagerIndex(srv.Ams, (value.(agt.FindWorkersRequest)).Id_manager)]).C_findWorkers <- srv.findWorkersRespond(value.(agt.FindWorkersRequest))
			default:
				fmt.Println("Error: bad request")
			}
		}
	}()

	go log.Fatal(s.ListenAndServe())
}

func (srv *Server) registerManager(am *agt.AgentManager) {
	fmt.Printf("Registering a manager. ID = %s\n", am.GetID())
	srv.Ams = append(srv.Ams, am)
}

func (srv *Server) registerWorker(aw *agt.AgentWorker) {
	fmt.Printf("Registering a worker. ID = %s\n", aw.GetID())
	srv.Aws = append(srv.Aws, aw)
}

func (srv *Server) findWorkersRespond(req agt.FindWorkersRequest) agt.FindWorkersResponse {
	var resp agt.FindWorkersResponse
	for _, v := range srv.Aws {
		if agt.ContainsHobby(v.Hobbies, req.Hobby) {
			resp.Workers = append(resp.Workers, v)
		}
	}
	return resp
}

func (srv *Server) GetManagers() []*agt.AgentManager {
	return srv.Ams
}

func (srv *Server) GetWorkers() []*agt.AgentWorker {
	return srv.Aws
}
