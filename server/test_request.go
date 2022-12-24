package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.utc.fr/pixelwar_ia04/pixelwar/agent"
)

// Méthode de la classe "Serveur" permettant de décoder une requête POST en JSON
func (*Server) decodeTestRequest(r *http.Request) (req testRequest, err error) {

	// Lecture de la requête POST
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)

	return
} // Retourne une requête de type "testRequest", et éventuellement une erreur

// Méthode de la classe "Server" permettant de décoder une requête POST en JSON
func (*Server) decodeRegisterAWRequest(r *http.Request) (req agent.RegisterAWRequest, err error) {

	// Lecture de la requête POST
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)

	return
} // Retourne une requête du type "registerAWRequest", et éventuellement une erreur

// Méthode de la classe "Server" permettant de décoder une requête POST en JSON
func (*Server) decodeRegisterAMRequest(r *http.Request) (req agent.RegisterAMRequest, err error) {

	// Lecture de la requête POST
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)

	return
} // Retourne une requête du type "registerAMRequest", et éventuellement une erreur

func (srv *Server) makeTestRequest(w http.ResponseWriter, r *http.Request) {

	// Verrouillage du serveur le temps de créer la requête
	srv.Lock()
	defer srv.Unlock()

	// Décodage de la requête POST
	req, err := srv.decodeTestRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Request OK. The value entered is: %s", req.Value)
	w.Write([]byte(msg))
}

func (srv *Server) registerAWRequest(w http.ResponseWriter, r *http.Request) {

	// Verrouillage du serveur le temps de traiter la requête
	srv.Lock()
	defer srv.Unlock()

	req, err := srv.decodeRegisterAWRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv.aws = append(srv.aws, req.Address)

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Request OK. The server successfuly registered the agent %s", (*req.Address).GetID())
	w.Write([]byte(msg))
}

func (srv *Server) registerAMRequest(w http.ResponseWriter, r *http.Request) {

	// Verrouillage du serveur le temps de traiter la requête
	srv.Lock()
	defer srv.Unlock()

	req, err := srv.decodeRegisterAMRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	srv.ams = append(srv.ams, req.Address)

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Request OK. The server successfuly registered the agent %s", (*req.Address).GetID())
	w.Write([]byte(msg))
}
