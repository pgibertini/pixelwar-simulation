package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/rs/cors"
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
	mux.HandleFunc("/canvas_old", srv.doCanvas)
	mux.HandleFunc("/get_diff", srv.doGetDiff)
	mux.HandleFunc("/canvas", srv.doCanvasDiff)

	// Création d'un serveur web
	s := &http.Server{
		Addr:           srv.address,
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20}

	log.Println("Listening on", srv.address)
	handler := cors.Default().Handler(mux)
    http.ListenAndServe(":5555", handler)
	go log.Fatal(s.ListenAndServe())
}
