package server

import (
	"log"
	"net/http"
	"time"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
)

func NewServer(id string, addr string) *Server {
	return &Server{
		identifier: id,
		address:    addr,
		// TODO: add a slices a Canvas (like in the vote API)
	}
}

func (srv *Server) Start() {
	// Multiplexage des différentes requêtes possibles
	mux := http.NewServeMux()
	mux.HandleFunc("/test_request", srv.makeTestRequest)
	mux.HandleFunc("/register_aw", srv.registerAWRequest)
	mux.HandleFunc("/register_am", srv.registerAMRequest)

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
	go log.Fatal(s.ListenAndServe())
}

func (srv *Server) GetManagers() []*agent.AgentManager {
	return srv.ams
}
