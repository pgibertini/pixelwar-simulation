package agent

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", true, "enable debug mode")
	// TODO: fix to have value passed by a flag
}

func NewServer(id string, addr string) *Server {
	return &Server{
		identifier: id,
		address:    addr,
		places:     make(map[string]*Place),
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
		Addr:           srv.address,
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Listening on", srv.address)
	go log.Fatal(s.ListenAndServe())
}
