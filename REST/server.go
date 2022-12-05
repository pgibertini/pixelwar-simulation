package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	sync.Mutex
	identifier string
	address    string
}

func NewServer(id string, addr string) *Server {
	return &Server{
		identifier: id,
		address:    addr,
	}

}

func (srv *Server) Start() {
	// Multiplexage des différentes requêtes possibles
	mux := http.NewServeMux()

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
