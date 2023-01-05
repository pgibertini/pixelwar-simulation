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

	go log.Fatal(s.ListenAndServe())
}
