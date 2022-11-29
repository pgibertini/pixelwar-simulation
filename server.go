package main

import "sync"

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

func (*Server) Start() {

}
