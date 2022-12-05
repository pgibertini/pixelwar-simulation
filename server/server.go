package server

import (
	"log"
	"net/http"
	"time"
)

func NewServer(id string, addr string) *Server {
	return &Server{
		identifier: id,
		address:    addr,
	}
}

func (srv *Server) Start() {
	// Multiplexage des différentes requêtes possibles
	mux := http.NewServeMux()
	mux.HandleFunc("/test_request", srv.makeTestRequest)

	// Création d'un serveur web
	s := &http.Server{
		Addr:           srv.address,
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Launching server server:", srv.address)
	go log.Fatal(s.ListenAndServe())
}
