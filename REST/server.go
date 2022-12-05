package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type testRequest struct {
	Value string `json:"value"`
}

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
	mux.HandleFunc("/test_request", srv.maketestRequest)

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

// Méthode de la classe "Serveur" permettant de décoder une requête POST en JSON
func (*Server) decodetestRequest(r *http.Request) (req testRequest, err error) {

	// Lecture de la requête POST
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)

	return
} // Retourne une requête de type "testRequest", et éventuellement une erreur

func (srv *Server) maketestRequest(w http.ResponseWriter, r *http.Request) {

	// Verrouillage du serveur le temps de créer la requête
	srv.Lock()
	defer srv.Unlock()

	// Décodage de la requête POST
	req, err := srv.decodetestRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Request OK. The value entered is: %s", req.Value)
	w.Write([]byte(msg))
	return
}

func main() {
	myServer := NewServer("TEST", "127.0.0.1:5555")
	go myServer.Start()

	fmt.Scanln()
}
